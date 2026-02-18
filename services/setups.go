package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundSetup struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	GetableResource[models.Setup, *BoundSetup]
	UpdatableResource[models.Setup, *BoundSetup]
}

func (b *BoundSetup) InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) *BoundSetup {
	if b == nil {
		b = &BoundSetup{}
	}
	b.nextDNSClient = nextDNSClient
	b.pathFmt = pathFmt
	b.pathArgs = pathArgs
	b.GetableResource.parent = b
	b.UpdatableResource.parent = b
	return b
}

func (b *BoundSetup) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundSetup) GetBoundPath() (string, error) {
	var path string
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

func (b *BoundSetup) GetError() error {
	return b.err
}

type SetupService struct {
	BindableResource[models.Setup, *BoundSetup]
}

func (c *NextDNSClient) Setup(profileId string) *BoundSetup {
	r := &SetupService{
		BindableResource: BindableResource[models.Setup, *BoundSetup]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/setup",
			pathArgs:      map[string]interface{}{"profileId": profileId},
		},
	}
	return r.Bind(profileId)
}
