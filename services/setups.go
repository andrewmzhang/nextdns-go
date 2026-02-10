package services

import "github.com/andrewmzhang/nextdns-go/models"

type BoundSetup struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	GetableResource[models.Setup]
	UpdatableResource[models.Setup]
	DeletableResource[models.Setup]
}

func (b *BoundSetup) SetBind(nextDNSClient *NextDNSClient, bindPath string) *BoundSetup {
	if b == nil {
		b = &BoundSetup{}
	}
	b.GetableResource.boundPath = bindPath
	b.GetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	b.DeletableResource.boundPath = bindPath
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type SetupService struct {
	// path string
	// ListableResource[models.Setup]
	// CreatableResource[models.Setup]
	BindableResource[models.Setup, *BoundSetup]
}

func (c *NextDNSClient) Setup(profileId string) *BoundSetup {
	r := &SetupService{
		BindableResource: BindableResource[models.Setup, *BoundSetup]{
			nextDNSClient: c,
			bindGeneratingFn: func(s string) string {
				return "/profiles/" + s + "/setup"
			},
		},
	}
	return r.Bind(profileId)
}
