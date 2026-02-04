package models

// Denylist represents the denylist of a profile.
//
//go:generate go run ./internal/gen/helpers Denylist
type Denylist struct {
	ID     string `json:"id,omitempty"`
	Active bool   `json:"active"`
}
