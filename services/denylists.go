package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundDenylist struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	ListGetableResource[models.Denylist, *BoundDenylist]
	UpdatableResource[models.Denylist, *BoundDenylist]
	DeletableResource[models.Denylist, *BoundDenylist]
}

func (b *BoundDenylist) InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) *BoundDenylist {
	if b == nil {
		b = &BoundDenylist{}
	}
	b.nextDNSClient = nextDNSClient
	b.pathFmt = pathFmt
	b.pathArgs = pathArgs
	b.err = err

	b.ListGetableResource.parent = b
	b.UpdatableResource.parent = b
	b.DeletableResource.parent = b
	return b
}

func (b *BoundDenylist) GetError() error {
	return b.err
}

func (b *BoundDenylist) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundDenylist) GetBoundPath() (string, error) {
	var path string
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

type DenylistService struct {
	// pathFmt string
	ListableResource[models.Denylist]
	CreatableResource[models.Denylist]
	BindableResource[models.Denylist, *BoundDenylist]
}

func (c *NextDNSClient) Denylists(profileId string) *DenylistService {
	return &DenylistService{
		ListableResource: ListableResource[models.Denylist]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/denylist",
			pathArgs: map[string]interface{}{
				"profileId": profileId,
			},
		},
		CreatableResource: CreatableResource[models.Denylist]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/denylist",
			pathArgs: map[string]interface{}{
				"profileId": profileId,
			},
		},
		BindableResource: BindableResource[models.Denylist, *BoundDenylist]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/denylist/{{ .id }}",
			pathArgs:      map[string]interface{}{"profileId": profileId},
		},
	}
}
