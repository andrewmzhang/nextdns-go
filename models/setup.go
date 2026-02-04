package models

import "encoding/json"

// Setup represents the setup settings.
type Setup struct {
	Ipv4     []string       `json:"ipv4"`
	Ipv6     []string       `json:"ipv6"`
	LinkedIP *SetupLinkedIP `json:"linkedIp"`
	Dnscrypt string         `json:"dnscrypt"`
}
type SetupLinkedIP struct {
	Servers     []string `json:"servers"`
	IP          string   `json:"ip"`
	Ddns        string   `json:"ddns"`
	UpdateToken string   `json:"updateToken"`
}

// MarshalJSON only ddns field is user settable
func (s SetupLinkedIP) MarshalJSON() ([]byte, error) {
	type output struct {
		Ddns string `json:"ddns"`
	}
	return json.Marshal(output{
		Ddns: s.Ddns,
	})
}
