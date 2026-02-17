package services

import (
	"fmt"

	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundRewrite struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	ListGetableResource[models.Rewrite, *BoundRewrite]
	DeletableResource[models.Rewrite, *BoundRewrite]
}

func (b *BoundRewrite) GetError() error {
	return b.err
}

func (b *BoundRewrite) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundRewrite) GetBoundPath() (string, error) {
	var path string
	fmt.Println("pathFmt", b.pathFmt, "pathArgs", b.pathArgs)
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

func (b *BoundRewrite) InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) *BoundRewrite {
	if b == nil {
		b = &BoundRewrite{}
	}
	b.nextDNSClient = nextDNSClient
	b.pathFmt = pathFmt
	b.pathArgs = pathArgs
	b.err = err
	b.ListGetableResource.parent = b
	b.DeletableResource.parent = b
	return b
}

type RewriteService struct {
	// pathFmt string
	ListableResource[models.Rewrite]
	CreatableResource[models.Rewrite]
	BindableResource[models.Rewrite, *BoundRewrite]
}

func (c *NextDNSClient) Rewrites(profileId string) *RewriteService {
	r := &RewriteService{
		ListableResource: ListableResource[models.Rewrite]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/rewrites",
			pathArgs: map[string]interface{}{
				"profileId": profileId,
			},
		},
		CreatableResource: CreatableResource[models.Rewrite]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/rewrites",
			pathArgs: map[string]interface{}{
				"profileId": profileId,
			},
		},
		BindableResource: BindableResource[models.Rewrite, *BoundRewrite]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/rewrites/{{ .id }}",
			pathArgs: map[string]interface{}{
				"profileId": profileId,
			},
		},
	}
	return r
}
