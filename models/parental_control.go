package models

import (
	"encoding/json"
)

//go:generate go run ./internal/gen/helpers ParentalControlRecreationInterval,ParentalControlRecreation,ParentalControlCategories,ParentalControlServices

// ParentalControlRecreationInterval represents the start and end time of a parental control recreation interval.
type ParentalControlRecreationInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// ParentalControlRecreationTimes represents the days and times of the week when the parental control is active.
type ParentalControlRecreationTimes struct {
	Monday    *ParentalControlRecreationInterval `json:"monday,omitempty"`
	Tuesday   *ParentalControlRecreationInterval `json:"tuesday,omitempty"`
	Wednesday *ParentalControlRecreationInterval `json:"wednesday,omitempty"`
	Thursday  *ParentalControlRecreationInterval `json:"thursday,omitempty"`
	Friday    *ParentalControlRecreationInterval `json:"friday,omitempty"`
	Saturday  *ParentalControlRecreationInterval `json:"saturday,omitempty"`
	Sunday    *ParentalControlRecreationInterval `json:"sunday,omitempty"`
}

// ParentalControlRecreation represents the parental control recreation of a profile.
type ParentalControlRecreation struct {
	Times    *ParentalControlRecreationTimes `json:"times"`
	Timezone string                          `json:"timezone"`
}

// ParentalControl represents the parental control settings of a profile.
type ParentalControl struct {
	Services              []*ParentalControlServices   `json:"services,omitempty"`
	Categories            []*ParentalControlCategories `json:"categories,omitempty"`
	Recreation            *ParentalControlRecreation   `json:"recreation,omitempty"`
	SafeSearch            *bool                        `json:"safeSearch"`
	YoutubeRestrictedMode *bool                        `json:"youtubeRestrictedMode"`
	BlockBypass           *bool                        `json:"blockBypass"`
}

// ParentalControlCategories represents the parental control categories of a profile.
//
//go:generate go run ./internal/gen/parentalcontrol
type ParentalControlCategoryID string
type ParentalControlCategories struct {
	ID         ParentalControlCategoryID `json:"id,omitempty"`
	Active     bool                      `json:"active"`
	Recreation bool                      `json:"recreation"`
}

type ParentalControlServiceID string

// ParentalControlServices represents the parental control services of a profile.
type ParentalControlServices struct {
	ID         ParentalControlServiceID `json:"id,omitempty"`
	Website    string                   `json:"website,omitempty"`
	Active     *bool                    `json:"active,omitzero"`
	Recreation *bool                    `json:"recreation,omitzero"`
}

func (p ParentalControlServices) MarshalJSON() ([]byte, error) {
	type Alias ParentalControlServices
	return json.Marshal(&struct {
		Alias
		Website *string `json:"website,omitempty"`
	}{
		Alias:   (Alias)(p),
		Website: nil,
	})
}
