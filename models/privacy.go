package models

import (
	"encoding/json"
	"time"
)

//go:generate go run ./internal/gen/helpers Privacy,PrivacyBlocklists,PrivacyNatives

// Privacy represents the privacy settings of a profile.
type Privacy struct {
	Blocklists        []*PrivacyBlocklists `json:"blocklists,omitzero"`
	Natives           []*PrivacyNatives    `json:"natives,omitzero"`
	DisguisedTrackers *bool                `json:"disguisedTrackers,omitzero"`
	AllowAffiliate    *bool                `json:"allowAffiliate,omitzero"`
}

type PrivacyBlocklistID string

// PrivacyBlocklists represents a privacy blocklist of a profile.
type PrivacyBlocklists struct {
	ID        PrivacyBlocklistID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Website   string             `json:"website,omitempty"`
	Entries   int                `json:"entries,omitempty"`
	UpdatedOn *time.Time         `json:"updatedOn,omitempty"`
}

func (p PrivacyBlocklists) MarshalJSON() ([]byte, error) {
	type Alias PrivacyBlocklists
	return json.Marshal(&struct {
		Alias
		Name      *string    `json:"name,omitempty"`
		Website   *string    `json:"website,omitempty"`
		Entries   *int       `json:"entries,omitempty"`
		UpdatedOn *time.Time `json:"updatedOn,omitempty"`
	}{
		Alias:     (Alias)(p),
		Name:      nil,
		Website:   nil,
		Entries:   nil,
		UpdatedOn: nil,
	})
}

//go:generate go run ./internal/gen/privacy
type PrivacyNativeID string

// PrivacyNatives represents a privacy native tracking protection of a profile.
type PrivacyNatives struct {
	ID PrivacyNativeID `json:"id"`
}
