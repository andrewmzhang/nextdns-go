package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/andrewmzhang/nextdns"
	"github.com/andrewmzhang/nextdns/models"
	"resty.dev/v3"
)

const (
	baseURL                    = "https://api.nextdns.io/"
	applicationJsonContentType = "application/json"
	nextdnsUserAgent           = "nextdns-go"
	linkedIPURL                = "https://link-ip.nextdns.io"
)

// NextDNSClient represents a NextDNS nextDNSClient.
type NextDNSClient struct {
	client      *resty.Client
	baseURL     *url.URL
	linkedIPURL *url.URL

	// Debug mode for the HTTP requests.
	Debug bool
}

// func (c *NextDNSClient) Allowlist(profileID string) []Resource[models.Allowlist] {
// 	return &Resource[models.Allowlist]{
// 		nextDNSClient: c,
// 		path:          "/profiles/" + profileID + "/allowlist",
// 	}
// }

//	func (c *NextDNSClient) analytics(profileID string) *Resource[models.Profile] {
//		return &Resource[models.ana]{
//			nextDNSClient: c,
//			path:          "/profiles/" + profileID + "/allowlist",
//		}
//	}
//
//	func (c *NextDNSClient) logs(profileID string) *Resource[models.Profile] {
//		return &Resource[models.SettingsLogs]{
//			nextDNSClient: c,
//			path:          "/profiles/" + profileID + "/allowlist",
//		}
//	}
func (c *NextDNSClient) settings(profileID string) *Resource[models.Settings] {
	return &Resource[models.Settings]{
		nextDNSClient: c,
		path:          "/profiles/" + profileID + "/settings",
	}
}
func (c *NextDNSClient) rewrites(profileID string) *Resource[models.Rewrite] {
	return &Resource[models.Rewrite]{
		nextDNSClient: c,
		path:          "/profiles/" + profileID + "/rewrites",
	}
}

// ClientOption is a function that can be used to customize the nextDNSClient.
type ClientOption func(c *NextDNSClient) error

// WithBaseURL sets the base URL of the NextDNS API.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *NextDNSClient) error {
		parsedURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		c.baseURL = parsedURL
		return nil
	}
}

// WithBaseURL sets the base URL of the NextDNS API.
func WithLinkedIPURL(linkedIPURL string) ClientOption {
	return func(c *NextDNSClient) error {
		parsedURL, err := url.Parse(linkedIPURL)
		if err != nil {
			return err
		}

		c.linkedIPURL = parsedURL
		return nil
	}
}

// WithAPIKey sets the API key to be used for requests.
func WithAPIKey(apiKey string) ClientOption {
	return func(c *NextDNSClient) error {
		if apiKey == "" {
			return nextdns.ErrEmptyAPIToken
		}
		c.client.SetHeader("X-api-key", apiKey)
		return nil
	}
}

// WithDebug enables debug mode.
func WithDebug() ClientOption {
	return func(c *NextDNSClient) error {
		c.Debug = true
		return nil
	}
}

// WithHTTPClient sets a resty HTTP nextDNSClient that can be used for requests.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *NextDNSClient) error {
		return nil
	}
}

// NewClient instantiates a new NextDNS nextDNSClient.
func NewClient(opts ...ClientOption) (*NextDNSClient, error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	nextDNSClient := &NextDNSClient{
		client: resty.New().
			SetBaseURL(baseURL.String()).
			SetHeader("Accept", applicationJsonContentType).
			SetHeader("Content-Type", applicationJsonContentType).
			OnError(func(req *resty.Request, err error) {
				if v, ok := err.(*resty.ResponseError); ok {
					// Do something with v.Response
					fmt.Printf("[ERROR] %s", v.Error())
				}
				// Log the error, increment a metric, etc...
			}),
		baseURL: baseURL,
	}

	for _, opt := range opts {
		err := opt(nextDNSClient)
		if err != nil {
			return nil, err
		}
	}

	// Initialize the services for the Profile.
	// nextDNSClient.Profiles = NewProfilesService(nextDNSClient)

	// // Initialize the services for the Allowlist and Denylist.
	// nextDNSClient.Allowlist = NewAllowlistService(nextDNSClient)
	// nextDNSClient.Denylist = NewDenylistService(nextDNSClient)
	//
	// // Initialize the services for the ParentalControl.
	// nextDNSClient.ParentalControl = NewParentalControlService(nextDNSClient)
	// nextDNSClient.ParentalControlServices = NewParentalControlServicesService(nextDNSClient)
	// nextDNSClient.ParentalControlCategories = NewParentalControlCategoriesService(nextDNSClient)
	//
	// // Initialize the services for the Privacy.
	// nextDNSClient.Privacy = NewPrivacyService(nextDNSClient)
	// nextDNSClient.PrivacyBlocklists = NewPrivacyBlocklistsService(nextDNSClient)
	// nextDNSClient.PrivacyNatives = NewPrivacyNativesService(nextDNSClient)
	//
	// // Initialize the services for the Settings.
	// nextDNSClient.Settings = NewSettingsService(nextDNSClient)
	// nextDNSClient.SettingsLogs = NewSettingsLogsService(nextDNSClient)
	// nextDNSClient.SettingsBlockPage = NewSettingsBlockPageService(nextDNSClient)
	// nextDNSClient.SettingsPerformance = NewSettingsPerformanceService(nextDNSClient)
	//
	// // Initialize the services for the Security.
	// nextDNSClient.Security = NewSecurityService(nextDNSClient)
	// nextDNSClient.SecurityTlds = NewSecurityTldsService(nextDNSClient)
	//
	// // Initialize the services for the Rewrites.
	// nextDNSClient.Rewrites = NewRewritesService(nextDNSClient)
	//
	// // Initialize the services for the Setup.
	// nextDNSClient.Setup = NewSetupService(nextDNSClient)
	// nextDNSClient.SetupLinkedIP = NewSetupLinkedIPService(nextDNSClient)

	return nextDNSClient, nil
}

func do[Resp any](
	ctx context.Context,
	c *NextDNSClient,
	method, path string,
	body any,
) (*Resp, error) {

	var resp *resty.Response
	var err error

	r := c.client.R().
		SetContext(ctx).
		SetResult(new(Resp)) // Resty will decode JSON into this

	if body != nil {
		r.SetBody(body)
	}

	switch method {
	case resty.MethodGet:
		resp, err = r.Get(path)
	case resty.MethodPost:
		resp, err = r.Post(path)
	case resty.MethodPut:
		resp, err = r.Put(path)
	case resty.MethodDelete:
		resp, err = r.Delete(path)
	default:
		return nil, fmt.Errorf("unsupported method %s", method)
	}

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("http error: %s", resp.Status())
	}

	result := resp.Result().(*Resp)
	return result, nil
}
