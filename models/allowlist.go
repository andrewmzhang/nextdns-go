package models

// Allowlist represents a single allowlist item
//
//go:generate go run ./internal/gen/helpers Allowlist
type Allowlist struct {
	ID     string `json:"id,omitempty"`
	Active bool   `json:"active"`
}
