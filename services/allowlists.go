package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundAllowlist struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	ListGetableResource[models.Allowlist, *BoundAllowlist]
	UpdatableResource[models.Allowlist, *BoundAllowlist]
	DeletableResource[models.Allowlist, *BoundAllowlist]
}

func (b *BoundAllowlist) InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) *BoundAllowlist {
	if b == nil {
		b = &BoundAllowlist{}
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

func (b *BoundAllowlist) GetError() error {
	return b.err
}

func (b *BoundAllowlist) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundAllowlist) GetBoundPath() (string, error) {
	var path string
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

type AllowlistService struct {
	// pathFmt string
	ListableResource[models.Allowlist]
	CreatableResource[models.Allowlist]
	BindableResource[models.Allowlist, *BoundAllowlist]
}

func (c *NextDNSClient) Allowlists(profileId string) *AllowlistService {
	return &AllowlistService{
		ListableResource: ListableResource[models.Allowlist]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/allowlist",
			pathArgs: map[string]interface{}{
				"profileId": profileId,
			},
		},
		CreatableResource: CreatableResource[models.Allowlist]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/allowlist",
			pathArgs: map[string]interface{}{
				"profileId": profileId,
			},
		},
		BindableResource: BindableResource[models.Allowlist, *BoundAllowlist]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/allowlist/{{ .id }}",
			pathArgs:      map[string]interface{}{"profileId": profileId},
		},
	}
}
