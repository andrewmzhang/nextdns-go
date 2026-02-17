package services

import (
	"encoding/hex"
	"strings"

	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundAllowlist struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	ListGetableResource[models.Allowlist]
	UpdatableResource[models.Allowlist]
	DeletableResource[models.Allowlist]
}

func (b *BoundAllowlist) SetBind(nextDNSClient *NextDNSClient, bindPath string, err error) *BoundAllowlist {
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
