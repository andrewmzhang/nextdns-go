package models

import (
	"encoding/json"
)

//go:generate go run ./internal/gen/helpers Profile

// Profile represents a NextDNS profile. Object structure closely resembles how io page is organized.
type Profile struct {
	// ID of the NextDNS Profile. This can be found in the profile's URL string and on the "Setup" tab -> "Endpoints"
	// -> "ID"
	ID string `json:"id,omitempty"`
	// Fingerprint of the NextDNS Profile TODO: Make this more clear
	Fingerprint string `json:"fingerprint,omitempty"`
	// Name of the NextDNS profile. Found on the top of the `io` page. Value is unique on a per-account basis
	Name string `json:"name,omitempty"`

	// Setup contains information found in the "Setup" tab on `io`
	Setup *Setup `json:"setup,omitempty"`
	// Security contains information found in the "Security" tab on `io`
	Security *Security `json:"security,omitempty"`
	// Privacy contains information found in the "Privacy" tab on `io`
	Privacy *Privacy `json:"privacy,omitempty"`
	// ParentalControl contains information found in the "Parental Control" tab on `io`
	ParentalControl *ParentalControl `json:"parentalControl,omitempty"`
	// Denylist contains a list of denied domains found on the "Denylist" tab on `io`
	Denylist []*Denylist `json:"denylist,omitzero"`
	// Allowlist contains a list of allowed domains found on the "Allowlist" tab on `io`
	Allowlist []*Allowlist `json:"allowlist,omitzero"`

	// TODO implement Analytics and Logs

	// Settings contains information found in the "Settings" tab on `io`, excluding the "Rewrites" section
	Settings *Settings `json:"settings,omitempty"`
	// TODO: Consider folding Rewrites under settings
	// Rewrites contains a list of dns domain rewrites found in the "Rewrites" section located near bottom of the
	// "Settings" tab on `io`
	Rewrites []*Rewrite `json:"rewrites,omitempty"`
}

func (p Profile) MarshalJSON() ([]byte, error) {
	type Alias Profile
	return json.Marshal(&struct {
		Alias
		ID          *string `json:"id,omitempty"`
		Fingerprint *string `json:"fingerprint,omitempty"`
		Setup       *string `json:"setup,omitempty"`
	}{
		Alias:       (Alias)(p),
		ID:          nil,
		Fingerprint: nil,
		Setup:       nil,
	})
}
