package services

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"resty.dev/v3"
)

// DataWrapper exists because NextDNS wraps API responses within a `data` field
type DataWrapper[T any] struct {
	Data T `json:"data"`
}

// DataListWrapper exists because NextDNS wraps API responses within a `data` field
type DataListWrapper[T any] struct {
	Data []T `json:"data"`
}

type Meta struct {
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type DataListMetaWrapper[T any] struct {
	Data []T  `json:"data"`
	Meta Meta `json:"meta"`
}

type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error (%d): %s", e.Status, e.Message)
}

func handleResponse[T any](resp *resty.Response, result *T) error {
	if resp.IsError() {
		// fallback if error body is not JSON / unknown format
		return &APIError{
			Status:  resp.StatusCode(),
			Message: resp.String(),
		}
	} else if resp.IsSuccess() {
		if resp.StatusCode() == http.StatusNoContent {
			return nil
		} else if result == nil {
			panic("Status code indicates response content, but no content available!")
		}
	}
	return nil
}

func validateNoEmptyStrings(data map[string]interface{}) error {
	for k, v := range data {
		if s, ok := v.(string); ok {
			if s == "" {
				return fmt.Errorf("field %q cannot be empty", k)
			}
		}
	}
	return nil
}

func renderPath(tmplStr string, data map[string]interface{}) (string, error) {
	if err := validateNoEmptyStrings(data); err != nil {
		return "", err
	}

	tmpl, err := template.New("tmpl").Option("missingkey=error").Parse(tmplStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Resource represents a NextDNS top-level construct, e.g. Profile
type ListableResource[T any] struct {
	nextDNSClient *NextDNSClient
	pathFmt       string
	pathArgs      map[string]interface{}
}

func (r *ListableResource[T]) List(
	ctx context.Context,
) ([]T, error) {
	path, err := renderPath(r.pathFmt, r.pathArgs)
	if err != nil {
		return nil, err
	}

	results := &struct {
		Data []T `json:"data"`
	}{}
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).SetResult(&results).Get(path)
	if err != nil {
		return nil, err
	}
	err = handleResponse(resp, results)
	return results.Data, err
}

type CreatableResource[T any] struct {
	nextDNSClient *NextDNSClient
	pathFmt       string
	pathArgs      map[string]interface{}
}

func (r *CreatableResource[T]) Create(ctx context.Context, payload T) (*T, error) {
	path, err := renderPath(r.pathFmt, r.pathArgs)
	if err != nil {
		return nil, err
	}
	result := &struct {
		Data T `json:"data"`
	}{}
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).SetBody(payload).SetResult(&result).Post(path)
	if err != nil {
		return nil, err
	}
	err = handleResponse(resp, result)
	return &result.Data, err
}

type BindableResource[T any, BoundResource HasBind[BoundResource]] struct {
	nextDNSClient *NextDNSClient
	pathFmt       string
	pathArgs      map[string]interface{}
}

func (r *BindableResource[T, BoundResource]) Bind(id string) BoundResource {
	var err error
	if id == "" {
		err = errors.New("id is required")
	}

	newPathArgs := map[string]interface{}{
		"id": id,
	}
	for k, v := range r.pathArgs {
		newPathArgs[k] = v
	}

	var v BoundResource
	return v.SetBind(r.nextDNSClient, r.pathFmt, newPathArgs, err)
}

type GetableResource[T any, Parent BoundResource] struct {
	parent Parent
}

func (r *GetableResource[T, Parent]) Get(ctx context.Context) (*T, error) {
	if r.parent.GetError() != nil {
		return nil, r.parent.GetError()
	}
	path, err := r.parent.GetBoundPath()
	if err != nil {
		return nil, err
	}
	fmt.Println("path:", path)
	var result DataWrapper[T]
	resp, err := r.parent.GetNextDNSClient().restyClient.R().SetContext(ctx).
		SetResult(&result).
		Get(path)
	if err != nil {
		return nil, err
	}
	err = handleResponse(resp, &result)
	return &result.Data, err
}

// ListGetableResource not all resource endpoints support get. The Get method under this type will call list on the parent
// resource and filter the results for desired item
type ListGetableResource[T any] struct {
	nextDNSClient *NextDNSClient
	boundPath     string
}

func (r *ListGetableResource[T]) Get(ctx context.Context) (*T, error) {
	results := &struct {
		Data []T `json:"data"`
	}{}
	parent := filepath.Dir(r.boundPath)
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).
		SetResult(&results).
		Get(parent)

	if err != nil {
		return nil, err
	}
	err = handleResponse(resp, &results)

	decodedId := filepath.Base(r.boundPath)
	if strings.HasPrefix(decodedId, "hex:") {
		decodedId = strings.TrimPrefix(decodedId, "hex:")
		bytes, err := hex.DecodeString(decodedId)
		if err != nil {
			return nil, err
		}
		decodedId = string(bytes)
	}

	for _, obj := range results.Data {
		// Marshal object to JSON
		b, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}

		// Unmarshal into a generic map
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, err
		}

		// Check ID
		if id, ok := m["id"].(string); ok && decodedId == id {
			return &obj, nil
		}
	}
	return nil, errors.New("Could not find resource with id: " + decodedId)
}

type UpdatableResource[T any, Parent BoundResource] struct {
	parent Parent
}

func (r *UpdatableResource[T, Parent]) Update(ctx context.Context, payload any) error {
	path, err := r.parent.GetBoundPath()
	if err != nil {
		return err
	}

	var result DataWrapper[T]
	resp, err := r.parent.GetNextDNSClient().restyClient.R().SetContext(ctx).
		SetBody(payload).
		Patch(path)

	err = handleResponse(resp, &result)
	return err
}

type DeletableResource[T any, Parent BoundResource] struct {
	parent Parent
}

func (r *DeletableResource[T, Parent]) Delete(ctx context.Context) error {
	path, err := r.parent.GetBoundPath()
	if err != nil {
		return err
	}
	var result DataWrapper[T]
	fmt.Println("path", path)
	resp, err := r.parent.GetNextDNSClient().restyClient.R().SetContext(ctx).
		Delete(path)
	if err != nil {
		return err
	}
	err = handleResponse(resp, &result)
	return err
}
