package models

//go:generate go run ./internal/gen/helpers Rewrite

// Rewrite represents a singular rewrite item under a profile.
type Rewrite struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Type    string `json:"type,omitempty"`
	Content string `json:"content"`
}
