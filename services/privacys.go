package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundPrivacy struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	GetableResource[models.Privacy]
	UpdatableResource[models.Privacy]
	// DeletableResource[models.Privacy]
}

func (b *BoundPrivacy) SetBind(nextDNSClient *NextDNSClient, bindPath string) *BoundPrivacy {
	if b == nil {
		b = &BoundPrivacy{}
	}
	b.GetableResource.boundPath = bindPath
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
		// ListableResource: ListableResource[models.Privacy]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		// CreatableResource: CreatableResource[models.Privacy]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
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
	// pathFmt string
	ListableResource[models.PrivacyBlocklists]
	// CreatableResource[models.Privacy]
	// BindableResource[models.Privacy, *BoundPrivacy]
}

func (c *NextDNSClient) PrivacyBlocklists() *PrivacyBlocklistsService {
	return &PrivacyBlocklistsService{
		ListableResource: ListableResource[models.PrivacyBlocklists]{
			nextDNSClient: c,
			pathFmt:       "/privacy/blocklists",
		},
		// CreatableResource: CreatableResource[models.Privacy]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		// BindableResource: BindableResource[models.Privacy, *BoundPrivacy]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/privacy"
		// 	},
		// },
	}
}

type PrivacyNativesService struct {
	// pathFmt string
	ListableResource[models.PrivacyNatives]
	// CreatableResource[models.Privacy]
	// BindableResource[models.Privacy, *BoundPrivacy]
}

func (c *NextDNSClient) PrivacyNatives() *PrivacyNativesService {
	return &PrivacyNativesService{
		ListableResource: ListableResource[models.PrivacyNatives]{
			nextDNSClient: c,
			pathFmt:       "/privacy/natives",
		},
		// CreatableResource: CreatableResource[models.Privacy]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		// BindableResource: BindableResource[models.Privacy, *BoundPrivacy]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/privacy"
		// 	},
		// },
	}
}
