package services

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

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

// Resource represents a NextDNS top-level construct, e.g. Profile
type Resource[T any] struct {
	nextDNSClient   *NextDNSClient
	path            string
	isSupportList   bool
	isSupportCreate bool
	isSupportBind   bool
}

func (r *Resource[T]) List(
	ctx context.Context,
) ([]T, error) {
	results := &struct {
		Data []T `json:"data"`
	}{}
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).SetResult(&results).Get(r.path)
	if err != nil {
		return nil, err
	}
	err = handleResponse(resp, results)
	return results.Data, err
}

func (r *Resource[T]) Create(ctx context.Context, payload T) (*T, error) {
	if !r.isSupportCreate {
		return nil, fmt.Errorf("create is not supported by this resource")
	}
	result := &struct {
		Data T `json:"data"`
	}{}
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).SetBody(payload).SetResult(&result).Post(r.path)
	if err != nil {
		return nil, err
	}
	err = handleResponse(resp, result)
	return &result.Data, err
}

func (r *Resource[T]) Bind(id string) *BoundResource[T] {
	return &BoundResource[T]{
		nextDNSClient: r.nextDNSClient,
		path:          r.path + "/" + fmt.Sprint(id),
	}
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

// Resource represents a NextDNS top-level construct, e.g. Profile
type ListableResource[T any] struct {
	nextDNSClient *NextDNSClient
	path          string
}

func (r *ListableResource[T]) List(
	ctx context.Context,
) ([]T, error) {
	results := &struct {
		Data []T `json:"data"`
	}{}
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).SetResult(&results).Get(r.path)
	if err != nil {
		return nil, err
	}
	err = handleResponse(resp, results)
	return results.Data, err
}

type CreatableResource[T any] struct {
	nextDNSClient *NextDNSClient
	path          string
}

func (r *CreatableResource[T]) Create(ctx context.Context, payload T) (*T, error) {
	result := &struct {
		Data T `json:"data"`
	}{}
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).SetBody(payload).SetResult(&result).Post(r.path)
	if err != nil {
		return nil, err
	}
	err = handleResponse(resp, result)
	return &result.Data, err
}

type BindableResource[T any, V HasBind[V]] struct {
	nextDNSClient    *NextDNSClient
	bindGeneratingFn func(string) string
}

func (r *BindableResource[T, V]) Bind(id string) V {
	if len(id) == 0 {
		panic("id is required")
	}

	var v V
	return v.SetBind(r.nextDNSClient, r.bindGeneratingFn(id))
}

type GetableResource[T any] struct {
	nextDNSClient *NextDNSClient
	boundPath     string
}

func (r *GetableResource[T]) Get(ctx context.Context) (*T, error) {

	var result DataWrapper[T]
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).
		SetResult(&result).
		Get(r.boundPath)
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

type UpdatableResource[T any] struct {
	nextDNSClient *NextDNSClient
	boundPath     string
}

func (r *UpdatableResource[T]) Update(ctx context.Context, payload any) error {
	var result DataWrapper[T]
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).
		SetBody(payload).
		Patch(r.boundPath)

	err = handleResponse(resp, &result)
	return err
}

type DeletableResource[T any] struct {
	nextDNSClient *NextDNSClient
	boundPath     string
}

func (r *DeletableResource[T]) Delete(ctx context.Context) error {
	var result DataWrapper[T]
	resp, err := r.nextDNSClient.restyClient.R().SetContext(ctx).
		Delete(r.boundPath)
	if err != nil {
		return err
	}
	err = handleResponse(resp, &result)
	return err
}
