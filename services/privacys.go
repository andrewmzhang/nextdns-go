package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundPrivacy struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	GetableResource[models.Privacy, *BoundPrivacy]
	UpdatableResource[models.Privacy, *BoundPrivacy]
}

func (b *BoundPrivacy) InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) *BoundPrivacy {
	if b == nil {
		b = &BoundPrivacy{}
	}
	b.nextDNSClient = nextDNSClient
	b.pathFmt = pathFmt
	b.pathArgs = pathArgs
	b.err = err
	b.GetableResource.parent = b
	b.UpdatableResource.parent = b
	return b
}

func (b *BoundPrivacy) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundPrivacy) GetBoundPath() (string, error) {
	var path string
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

func (b *BoundPrivacy) GetError() error {
	return b.err
}

type PrivacyService struct {
	BindableResource[models.Privacy, *BoundPrivacy]
}

func (c *NextDNSClient) Privacy(profileId string) *BoundPrivacy {
	r := &PrivacyService{
		BindableResource: BindableResource[models.Privacy, *BoundPrivacy]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/privacy",
			pathArgs:      map[string]interface{}{"profileId": profileId},
		},
	}
	return r.Bind(profileId)
}

type PrivacyBlocklistsService struct {
	ListableResource[models.PrivacyBlocklists]
}

func (c *NextDNSClient) PrivacyBlocklists() *PrivacyBlocklistsService {
	return &PrivacyBlocklistsService{
		ListableResource: ListableResource[models.PrivacyBlocklists]{
			nextDNSClient: c,
			pathFmt:       "/privacy/blocklists",
		},
	}
}

type PrivacyNativesService struct {
	// pathFmt string
	ListableResource[models.PrivacyNatives]
}

func (c *NextDNSClient) PrivacyNatives() *PrivacyNativesService {
	return &PrivacyNativesService{
		ListableResource: ListableResource[models.PrivacyNatives]{
			nextDNSClient: c,
			pathFmt:       "/privacy/natives",
		},
	}
}
