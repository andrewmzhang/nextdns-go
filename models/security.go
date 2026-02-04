package models

import (
	"encoding/json"
)

//go:generate go run ./internal/gen/helpers Security,SecurityTlds

// Security represents the security settings of a profile.
type Security struct {
	ThreatIntelligenceFeeds bool            `json:"threatIntelligenceFeeds"`
	AiThreatDetection       bool            `json:"aiThreatDetection"`
	GoogleSafeBrowsing      bool            `json:"googleSafeBrowsing"`
	Cryptojacking           bool            `json:"cryptojacking"`
	DNSRebinding            bool            `json:"dnsRebinding"`
	IdnHomographs           bool            `json:"idnHomographs"`
	Typosquatting           bool            `json:"typosquatting"`
	Dga                     bool            `json:"dga"`
	Nrd                     bool            `json:"nrd"`
	DDNS                    bool            `json:"ddns"`
	Parking                 bool            `json:"parking"`
	Csam                    bool            `json:"csam"`
	BlockedTLDs             []*SecurityTlds `json:"tlds,omitzero"`
}

type SecurityTldID string

// SecurityTlds represents the security TLDs of a profile.
//
//go:generate go run ./internal/gen/securitytld
type SecurityTlds struct {
	ID       SecurityTldID `json:"id"`
	Spamhaus int           `json:"spamhaus,omitempty"` // field only exists when querying <baseurl>/security/tlds
}

func (s SecurityTlds) MarshalJSON() ([]byte, error) {
	type Alias SecurityTlds
	return json.Marshal(&struct {
		Alias
		Spamhaus *int `json:"spamhaus,omitempty"`
	}{
		Alias:    (Alias)(s),
		Spamhaus: nil,
	})
}
