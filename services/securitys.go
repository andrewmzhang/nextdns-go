package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundSecurity struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	GetableResource[models.Security]
	UpdatableResource[models.Security]
	DeletableResource[models.Security]
}

func (b *BoundSecurity) SetBind(nextDNSClient *NextDNSClient, bindPath string) *BoundSecurity {
	if b == nil {
		b = &BoundSecurity{}
	}
	b.GetableResource.boundPath = bindPath
	b.GetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	b.DeletableResource.boundPath = bindPath
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type SecurityService struct {
	// path string
	// ListableResource[models.Security]
	// CreatableResource[models.Security]
	BindableResource[models.Security, *BoundSecurity]
}

func (c *NextDNSClient) Security(profileId string) *BoundSecurity {
	r := &SecurityService{
		// ListableResource: ListableResource[models.Security]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		// CreatableResource: CreatableResource[models.Security]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		BindableResource: BindableResource[models.Security, *BoundSecurity]{
			nextDNSClient: c,
			bindGeneratingFn: func(s string) string {
				return "/profiles/" + s + "/security"
			},
		},
	}
	return r.Bind(profileId)
}

type SecurityTldsService struct {
	// path string
	ListableResource[models.SecurityTlds]
	// CreatableResource[models.Security]
	// BindableResource[models.Security, *BoundSecurity]
}

func (c *NextDNSClient) SecurityTlds() *SecurityTldsService {
	return &SecurityTldsService{
		ListableResource: ListableResource[models.SecurityTlds]{
			nextDNSClient: c,
			path:          "/security/tlds",
		},
		// CreatableResource: CreatableResource[models.Security]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		// BindableResource: BindableResource[models.Security, *BoundSecurity]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/security"
		// 	},
		// },
	}
}
