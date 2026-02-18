package services

import (
	"fmt"

	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundProfile struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	GetableResource[models.Profile, *BoundProfile]
	UpdatableResource[models.Profile, *BoundProfile]
	DeletableResource[models.Profile, *BoundProfile]
}

func (b *BoundProfile) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundProfile) GetBoundPath() (string, error) {
	var path string
	fmt.Println("pathFmt", b.pathFmt, "pathArgs", b.pathArgs)
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

func (b *BoundProfile) GetError() error {
	return b.err
}

func (b *BoundProfile) InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) *BoundProfile {
	if b == nil {
		b = &BoundProfile{}
	}
	b.nextDNSClient = nextDNSClient
	b.pathFmt = pathFmt
	b.pathArgs = pathArgs
	b.err = err
	b.GetableResource.parent = b
	b.UpdatableResource.parent = b
	b.DeletableResource.parent = b
	return b
}

type ProfileService struct {
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
			pathFmt:       "/profiles/{{ .id }}",
		},
	}
}
