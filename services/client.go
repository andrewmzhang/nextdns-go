package services

import (
	"net/http"
	"net/url"

	"resty.dev/v3"
)

// NextDNSClient represents a NextDNS nextDNSClient.
type NextDNSClient struct {
	restyClient *resty.Client
	baseURL     *url.URL
	linkedIPURL *url.URL

	// Debug mode for the HTTP requests.
	Debug bool
}

// ClientOption is a function that can be used to customize the nextDNSClient.
type ClientOption func(c *NextDNSClient) error

func NewNextDNSClient(client *resty.Client, baseURL *url.URL) *NextDNSClient {
	return &NextDNSClient{restyClient: client, baseURL: baseURL}
}

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
			return ErrEmptyAPIToken
		}
		c.restyClient.SetHeader("X-api-key", apiKey)
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

// WithHTTPClient sets a resty HTTP NextDNSClient that can be used for requests.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *NextDNSClient) error {
		return nil
	}
}
