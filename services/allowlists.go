package services

import (
	"encoding/hex"
	"strings"

	"github.com/andrewmzhang/nextdns/models"
)

type BoundAllowlist struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	ListGetableResource[models.Allowlist]
	UpdatableResource[models.Allowlist]
	DeletableResource[models.Allowlist]
}

func (b *BoundAllowlist) SetBind(nextDNSClient *NextDNSClient, bindPath string) *BoundAllowlist {
	if b == nil {
		b = &BoundAllowlist{}
	}
	b.ListGetableResource.boundPath = bindPath
	b.ListGetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	b.DeletableResource.boundPath = bindPath
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type AllowlistService struct {
	// path string
	ListableResource[models.Allowlist]
	CreatableResource[models.Allowlist]
	BindableResource[models.Allowlist, *BoundAllowlist]
}

func (c *NextDNSClient) Allowlists(profileId string) *AllowlistService {
	return &AllowlistService{
		ListableResource: ListableResource[models.Allowlist]{
			nextDNSClient: c,
			path:          "/profiles/" + profileId + "/allowlist",
		},
		CreatableResource: CreatableResource[models.Allowlist]{
			nextDNSClient: c,
			path:          "/profiles/" + profileId + "/allowlist",
		},
		BindableResource: BindableResource[models.Allowlist, *BoundAllowlist]{
			nextDNSClient: c,
			bindGeneratingFn: func(s string) string {
				if strings.HasPrefix(s, "hex:") {
					return "/profiles/" + profileId + "/allowlist/" + s
				}
				dst := make([]byte, hex.EncodedLen(len(s)))
				hex.Encode(dst, []byte(s))
				return "/profiles/" + profileId + "/allowlist/hex:" + string(dst)
			},
		},
	}
}
