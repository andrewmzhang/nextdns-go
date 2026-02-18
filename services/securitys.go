package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundSecurity struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	GetableResource[models.Security, *BoundSecurity]
	UpdatableResource[models.Security, *BoundSecurity]
}

func (b *BoundSecurity) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundSecurity) GetBoundPath() (string, error) {
	var path string
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

func (b *BoundSecurity) GetError() error {
	return b.err
}

func (b *BoundSecurity) InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) *BoundSecurity {
	if b == nil {
		b = &BoundSecurity{}
	}
	b.nextDNSClient = nextDNSClient
	b.pathFmt = pathFmt
	b.pathArgs = pathArgs
	b.err = err
	b.GetableResource.parent = b
	b.UpdatableResource.parent = b
	return b
}

type SecurityService struct {
	BindableResource[models.Security, *BoundSecurity]
}

func (c *NextDNSClient) Security(profileId string) *BoundSecurity {
	r := &SecurityService{
		BindableResource: BindableResource[models.Security, *BoundSecurity]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/security",
			pathArgs:      map[string]interface{}{"profileId": profileId},
		},
	}
	return r.Bind(profileId)
}

type SecurityTldsService struct {
	ListableResource[models.SecurityTlds]
}

func (c *NextDNSClient) SecurityTlds() *SecurityTldsService {
	return &SecurityTldsService{
		ListableResource: ListableResource[models.SecurityTlds]{
			nextDNSClient: c,
			pathFmt:       "/security/tlds",
		},
	}
}
