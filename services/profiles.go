package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundProfile struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	GetableResource[models.Profile, *BoundProfile]
	UpdatableResource[models.Profile]
	DeletableResource[models.Profile]
}

func (b *BoundProfile) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundProfile) GetBoundPath() (string, error) {
	var path string
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

func (b *BoundProfile) GetError() error {
	return b.err
}

func (b *BoundProfile) SetBind(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}) *BoundProfile {
	if b == nil {
		b = &BoundProfile{}
	}
	b.GetableResource.parent = b
	b.GetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = pathFmt
	b.UpdatableResource.nextDNSClient = nextDNSClient
	b.DeletableResource.boundPath = pathFmt
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type ProfileService struct {
	// pathFmt string
	ListableResource[models.Profile]
	CreatableResource[models.Profile]
	BindableResource[models.Profile, *BoundProfile]
}

func (c *NextDNSClient) Profiles() *ProfileService {
	return &ProfileService{
		ListableResource: ListableResource[models.Profile]{
			nextDNSClient: c,
			pathFmt:       "/profiles",
		},
		CreatableResource: CreatableResource[models.Profile]{
			nextDNSClient: c,
			pathFmt:       "/profiles",
		},
		BindableResource: BindableResource[models.Profile, *BoundProfile]{
			nextDNSClient: c,
			bindGeneratingFn: func(s string) string {
				return "/profiles/" + s
			},
		},
	}
}
