package services

import (
	"encoding/hex"
	"strings"

	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundDenylist struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	ListGetableResource[models.Denylist]
	UpdatableResource[models.Denylist]
	DeletableResource[models.Denylist]
}

func (b *BoundDenylist) SetBind(nextDNSClient *NextDNSClient, bindPath string, err error) *BoundDenylist {
	if b == nil {
		b = &BoundDenylist{}
	}
	b.ListGetableResource.boundPath = bindPath
	b.ListGetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	b.DeletableResource.boundPath = bindPath
	b.DeletableResource.nextDNSClient = nextDNSClient
	return b
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
			bindGeneratingFn: func(s string) string {
				if strings.HasPrefix(s, "hex:") {
					return "/profiles/" + profileId + "/denylist/" + s
				}
				dst := make([]byte, hex.EncodedLen(len(s)))
				hex.Encode(dst, []byte(s))
				return "/profiles/" + profileId + "/denylist/hex:" + string(dst)
			},
		},
	}
}
