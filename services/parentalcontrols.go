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

func (b *BoundParentalControl) SetBind(nextDNSClient *NextDNSClient, bindPath string) *BoundParentalControl {
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
	// path string
	// ListableResource[models.ParentalControl]
	// CreatableResource[models.ParentalControl]
	BindableResource[models.ParentalControl, *BoundParentalControl]
}

func (c *NextDNSClient) ParentalControl(profileId string) *BoundParentalControl {
	r := &ParentalControlService{
		// ListableResource: ListableResource[models.ParentalControl]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		// CreatableResource: CreatableResource[models.ParentalControl]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
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
	// path string
	ListableResource[models.ParentalControlServices]
	// CreatableResource[models.ParentalControl]
	// BindableResource[models.ParentalControl, *BoundParentalControl]
}

func (c *NextDNSClient) ParentalControlServices() *ParentalControlServicesService {
	return &ParentalControlServicesService{
		ListableResource: ListableResource[models.ParentalControlServices]{
			nextDNSClient: c,
			path:          "/parentalcontrol/services",
		},
		// CreatableResource: CreatableResource[models.ParentalControl]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		// BindableResource: BindableResource[models.ParentalControl, *BoundParentalControl]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/parentalcontrol"
		// 	},
		// },
	}
}

type ParentalControlCategoriesService struct {
	// path string
	ListableResource[models.ParentalControlCategories]
	// CreatableResource[models.ParentalControl]
	// BindableResource[models.ParentalControl, *BoundParentalControl]
}

func (c *NextDNSClient) ParentalControlCategories() *ParentalControlCategoriesService {
	return &ParentalControlCategoriesService{
		ListableResource: ListableResource[models.ParentalControlCategories]{
			nextDNSClient: c,
			path:          "/parentalcontrol/categories",
		},
		// CreatableResource: CreatableResource[models.ParentalControl]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// },
		// BindableResource: BindableResource[models.ParentalControl, *BoundParentalControl]{
		// 	nextDNSClient: c,
		// 	path:          "/profiles",
		// 	bindGeneratingFn: func(s string) string {
		// 		return "/profiles/" + s + "/parentalcontrol"
		// 	},
		// },
	}
}
