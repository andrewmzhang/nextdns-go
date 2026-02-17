package services

import (
	"github.com/andrewmzhang/nextdns-go/models"
)

type BoundParentalControl struct {
	boundPath     string
	nextDNSClient *NextDNSClient
	GetableResource[models.ParentalControl]
	UpdatableResource[models.ParentalControl]
	// DeletableResource[models.ParentalControl]
}

func (b *BoundParentalControl) SetBind(nextDNSClient *NextDNSClient, bindPath string, err error) *BoundParentalControl {
	if b == nil {
		b = &BoundParentalControl{}
	}
	b.GetableResource.boundPath = bindPath
	b.GetableResource.nextDNSClient = nextDNSClient
	b.UpdatableResource.boundPath = bindPath
	b.UpdatableResource.nextDNSClient = nextDNSClient
	// b.DeletableResource.boundPath = bindPath
	// b.DeletableResource.nextDNSClient = nextDNSClient
	return b
}

type ParentalControlService struct {
	// pathFmt string
	// ListableResource[models.ParentalControl]
	// CreatableResource[models.ParentalControl]
	BindableResource[models.ParentalControl, *BoundParentalControl]
}

func (c *NextDNSClient) ParentalControl(profileId string) *BoundParentalControl {
	r := &ParentalControlService{
		// ListableResource: ListableResource[models.ParentalControl]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		// CreatableResource: CreatableResource[models.ParentalControl]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		BindableResource: BindableResource[models.ParentalControl, *BoundParentalControl]{
			nextDNSClient: c,
			bindGeneratingFn: func(s string) string {
				return "/profiles/" + s + "/parentalcontrol"
			},
		},
	}
	return r.Bind(profileId)
}

type ParentalControlServicesService struct {
	// pathFmt string
	ListableResource[models.ParentalControlServices]
	// CreatableResource[models.ParentalControl]
	// BindableResource[models.ParentalControl, *BoundParentalControl]
}

func (c *NextDNSClient) ParentalControlServices() *ParentalControlServicesService {
	return &ParentalControlServicesService{
		ListableResource: ListableResource[models.ParentalControlServices]{
			nextDNSClient: c,
			pathFmt:       "/parentalcontrol/services",
		},
		// CreatableResource: CreatableResource[models.ParentalControl]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		// BindableResource: BindableResource[models.ParentalControl, *BoundParentalControl]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/parentalcontrol"
		// 	},
		// },
	}
}

type ParentalControlCategoriesService struct {
	// pathFmt string
	ListableResource[models.ParentalControlCategories]
	// CreatableResource[models.ParentalControl]
	// BindableResource[models.ParentalControl, *BoundParentalControl]
}

func (c *NextDNSClient) ParentalControlCategories() *ParentalControlCategoriesService {
	return &ParentalControlCategoriesService{
		ListableResource: ListableResource[models.ParentalControlCategories]{
			nextDNSClient: c,
			pathFmt:       "/parentalcontrol/categories",
		},
		// CreatableResource: CreatableResource[models.ParentalControl]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// },
		// BindableResource: BindableResource[models.ParentalControl, *BoundParentalControl]{
		// 	nextDNSClient: c,
		// 	pathFmt:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/parentalcontrol"
		// 	},
		// },
	}
}
