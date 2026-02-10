package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundProfile struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	GetableResource[models.Profile]
	UpdatableResource[models.Profile]
	DeletableResource[models.Profile]
}

func (b *BoundProfile) SetBind(nextDNSClient *NextDNSClient, bindPath string) *BoundProfile {
	if b == nil {
		b = &BoundProfile{}
	}
	b.GetableResource.boundPath = bindPath
	b.GetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	b.DeletableResource.boundPath = bindPath
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type ProfileService struct {
	// path string
	ListableResource[models.Profile]
	CreatableResource[models.Profile]
	BindableResource[models.Profile, *BoundProfile]
}

func (c *NextDNSClient) Profiles() *ProfileService {
	return &ProfileService{
		ListableResource: ListableResource[models.Profile]{
			nextDNSClient: c,
			path:          "/profiles",
		},
		CreatableResource: CreatableResource[models.Profile]{
			nextDNSClient: c,
			path:          "/profiles",
		},
		BindableResource: BindableResource[models.Profile, *BoundProfile]{
			nextDNSClient: c,
			bindGeneratingFn: func(s string) string {
				return "/profiles/" + s
			},
		},
	}
}
