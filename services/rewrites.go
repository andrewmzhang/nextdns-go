package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundRewrite struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	// GetableResource[models.Rewrite]
	// UpdatableResource[models.Rewrite]
	ListGetableResource[models.Rewrite]
	DeletableResource[models.Rewrite]
}

func (b *BoundRewrite) SetBind(nextDNSClient *NextDNSClient, bindPath string) *BoundRewrite {
	if b == nil {
		b = &BoundRewrite{}
	}
	// b.GetableResource.boundPath = bindPath
	// b.GetableResource.nextDNSClient = nextDNSClient
	// b.UpdatableResource.boundPath = bindPath
	// b.UpdatableResource.nextDNSClient = nextDNSClient
	b.ListGetableResource.nextDNSClient = nextDNSClient
	b.ListGetableResource.boundPath = bindPath
	b.DeletableResource.boundPath = bindPath
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type RewriteService struct {
	// path string
	ListableResource[models.Rewrite]
	CreatableResource[models.Rewrite]
	BindableResource[models.Rewrite, *BoundRewrite]
}

func (c *NextDNSClient) Rewrites(profileId string) *RewriteService {
	r := &RewriteService{
		ListableResource: ListableResource[models.Rewrite]{
			nextDNSClient: c,
			path:          "/profiles/" + profileId + "/rewrites",
		},
		CreatableResource: CreatableResource[models.Rewrite]{
			nextDNSClient: c,
			path:          "/profiles/" + profileId + "/rewrites",
		},
		BindableResource: BindableResource[models.Rewrite, *BoundRewrite]{
			nextDNSClient: c,
			bindGeneratingFn: func(s string) string {
				return "/profiles/" + profileId + "/rewrites/" + s
			},
		},
	}
	return r
}
