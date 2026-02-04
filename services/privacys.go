package services

import "github.com/andrewmzhang/nextdns/models"

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
	// path string
	// ListableResource[models.Privacy]
	// CreatableResource[models.Privacy]
	BindableResource[models.Privacy, *BoundPrivacy]
}

func (c *NextDNSClient) Privacy(profileId string) *BoundPrivacy {
	r := &PrivacyService{
		// ListableResource: ListableResource[models.Privacy]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		// CreatableResource: CreatableResource[models.Privacy]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
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
	// path string
	ListableResource[models.PrivacyBlocklists]
	// CreatableResource[models.Privacy]
	// BindableResource[models.Privacy, *BoundPrivacy]
}

func (c *NextDNSClient) PrivacyBlocklists() *PrivacyBlocklistsService {
	return &PrivacyBlocklistsService{
		ListableResource: ListableResource[models.PrivacyBlocklists]{
			nextDNSClient: c,
			path:          "/privacy/blocklists",
		},
		// CreatableResource: CreatableResource[models.Privacy]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		// BindableResource: BindableResource[models.Privacy, *BoundPrivacy]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/privacy"
		// 	},
		// },
	}
}

type PrivacyNativesService struct {
	// path string
	ListableResource[models.PrivacyNatives]
	// CreatableResource[models.Privacy]
	// BindableResource[models.Privacy, *BoundPrivacy]
}

func (c *NextDNSClient) PrivacyNatives() *PrivacyNativesService {
	return &PrivacyNativesService{
		ListableResource: ListableResource[models.PrivacyNatives]{
			nextDNSClient: c,
			path:          "/privacy/natives",
		},
		// CreatableResource: CreatableResource[models.Privacy]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		// BindableResource: BindableResource[models.Privacy, *BoundPrivacy]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/privacy"
		// 	},
		// },
	}
}
