package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundSetup struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	GetableResource[models.Setup, *BoundSetup]
	UpdatableResource[models.Setup]
	DeletableResource[models.Setup]
}

func (b *BoundSetup) GetError() error {
	// Todo Implement this
	return nil
}

func (b *BoundSetup) SetBind(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}) T {
	if b == nil {
		b = &BoundSetup{}
	}
	b.GetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	b.DeletableResource.boundPath = bindPath
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type SetupService struct {
	// pathFmt string
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
