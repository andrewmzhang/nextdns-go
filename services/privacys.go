package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundPrivacy struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	err           error
	GetableResource[models.Privacy, *BoundPrivacy]
	UpdatableResource[models.Privacy]
	// DeletableResource[models.Privacy]
}

func (b *BoundPrivacy) GetError() error {
	return b.err
}

func (b *BoundPrivacy) SetBind(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}) T {
	if b == nil {
		b = &BoundPrivacy{}
	}
	b.GetableResource.parent = b
	b.GetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	// b.DeletableResource.boundPath = bindPath
	// b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type PrivacyService struct {
	// pathFmt string
	// ListableResource[models.Privacy]
	// CreatableResource[models.Privacy]
	BindableResource[models.Privacy, *BoundPrivacy]
}

func (c *NextDNSClient) Privacy(profileId string) *BoundPrivacy {
	r := &PrivacyService{
		BindableResource: BindableResource[models.Privacy, *BoundPrivacy]{
			nextDNSClient: c,
			bindGeneratingFn: func(s string) string {
				return "/profiles/" + s + "/privacy"
			},
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
