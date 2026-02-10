package nextdns

import (
	"fmt"
	"net/url"

	"github.com/andrewmzhang/nextdns-go/services"
	"resty.dev/v3"
)

const (
	baseURL                    = "https://api.nextdns.io/"
	applicationJsonContentType = "application/json"
	nextdnsUserAgent           = "nextdns-go"
	linkedIPURL                = "https://link-ip.nextdns.io"
)

// NewClient instantiates a new NextDNS nextDNSClient.
func NewClient(opts ...services.ClientOption) (*services.NextDNSClient, error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	nextDNSClient := services.NewNextDNSClient(
		resty.New().
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
		baseURL,
	)

	for _, opt := range opts {
		err := opt(nextDNSClient)
		if err != nil {
			return nil, err
		}
	}

	return nextDNSClient, nil
}
