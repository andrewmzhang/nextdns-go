package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundSecurity struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	err           error
	GetableResource[models.Security, *BoundSecurity]
	UpdatableResource[models.Security]
	DeletableResource[models.Security]
}

func (b *BoundSecurity) GetError() error {
	return b.err
}

func (b *BoundSecurity) SetBind(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}) T {
	if b == nil {
		b = &BoundSecurity{}
	}
	b.GetableResource.parent = b
	b.GetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	b.DeletableResource.boundPath = bindPath
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type SecurityService struct {
	// pathFmt string
	// ListableResource[models.Security]
	// CreatableResource[models.Security]
	BindableResource[models.Security, *BoundSecurity]
}

func (c *NextDNSClient) Security(profileId string) *BoundSecurity {
	r := &SecurityService{
		// ListableResource: ListableResource[models.Security]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		// CreatableResource: CreatableResource[models.Security]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
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
	// pathFmt string
	ListableResource[models.SecurityTlds]
	// CreatableResource[models.Security]
	// BindableResource[models.Security, *BoundSecurity]
}

func (c *NextDNSClient) SecurityTlds() *SecurityTldsService {
	return &SecurityTldsService{
		ListableResource: ListableResource[models.SecurityTlds]{
			nextDNSClient: c,
			pathFmt:       "/security/tlds",
		},
		// CreatableResource: CreatableResource[models.Security]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		// BindableResource: BindableResource[models.Security, *BoundSecurity]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/security"
		// 	},
		// },
	}
}
