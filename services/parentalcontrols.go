package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundParentalControl struct {
	err           error
	pathFmt       string
	pathArgs      map[string]interface{}
	nextDNSClient *NextDNSClient
	GetableResource[models.ParentalControl, *BoundParentalControl]
	UpdatableResource[models.ParentalControl, *BoundParentalControl]
}

func (b *BoundParentalControl) InitBoundResource(nextDNSClient *NextDNSClient, pathFmt string, pathArgs map[string]interface{}, err error) *BoundParentalControl {
	if b == nil {
		b = &BoundParentalControl{}
	}
	b.nextDNSClient = nextDNSClient
	b.pathFmt = pathFmt
	b.pathArgs = pathArgs
	b.err = err
	b.GetableResource.parent = b
	b.UpdatableResource.parent = b
	return b
}

func (b *BoundParentalControl) GetNextDNSClient() *NextDNSClient {
	return b.nextDNSClient
}

func (b *BoundParentalControl) GetBoundPath() (string, error) {
	var path string
	path, b.err = renderPath(b.pathFmt, b.pathArgs)
	return path, b.err
}

func (b *BoundParentalControl) GetError() error {
	return b.err
}

type ParentalControlService struct {
	BindableResource[models.ParentalControl, *BoundParentalControl]
}

func (c *NextDNSClient) ParentalControl(profileId string) *BoundParentalControl {
	r := &ParentalControlService{
		BindableResource: BindableResource[models.ParentalControl, *BoundParentalControl]{
			nextDNSClient: c,
			pathFmt:       "/profiles/{{ .profileId }}/parentalcontrol",
			pathArgs:      map[string]interface{}{"profileId": profileId},
		},
	}
	return r.Bind(profileId)
}

type ParentalControlServicesService struct {
	ListableResource[models.ParentalControlServices]
}

func (c *NextDNSClient) ParentalControlServices() *ParentalControlServicesService {
	return &ParentalControlServicesService{
		ListableResource: ListableResource[models.ParentalControlServices]{
			nextDNSClient: c,
			pathFmt:       "/parentalcontrol/services",
		},
	}
}

type ParentalControlCategoriesService struct {
	ListableResource[models.ParentalControlCategories]
}

func (c *NextDNSClient) ParentalControlCategories() *ParentalControlCategoriesService {
	return &ParentalControlCategoriesService{
		ListableResource: ListableResource[models.ParentalControlCategories]{
			nextDNSClient: c,
			pathFmt:       "/parentalcontrol/categories",
		},
	}
}
