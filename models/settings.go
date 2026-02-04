package models

//go:generate go run ./internal/gen/helpers Settings,SettingsBlockPage,SettingsLogsDrop,SettingsLogs,SettingsPerformance

// Settings represents the settings of a profile.
type Settings struct {
	Logs        *SettingsLogs        `json:"logs,omitempty"`
	BlockPage   *SettingsBlockPage   `json:"blockPage,omitempty"`
	Performance *SettingsPerformance `json:"performance,omitempty"`
	Web3        bool                 `json:"web3"`
}

// SettingsBlockPage represents the settings block page of a profile.
type SettingsBlockPage struct {
	Enabled *bool `json:"enabled,omitzero"`
}

// SettingsLogsDrop represents the settings logs privacy adjustments of a profile.
type SettingsLogsDrop struct {
	IP     *bool `json:"ip,omitzero"`
	Domain *bool `json:"domain,omitzero"`
}

type SettingsLocationID string

const (
	SettingsLocationUnitedStates  SettingsLocationID = "us"
	SettingsLocationEuropeanUnion SettingsLocationID = "eu"
	SettingsLocationGreatBritain  SettingsLocationID = "gb"
	SettingsLocationSwitzerland   SettingsLocationID = "ch"
)

type SettingsRetentionID int

const (
	SettingsRetention1Hour   SettingsRetentionID = 60 * 60
	SettingsRetention6Hours  SettingsRetentionID = 6 * SettingsRetention1Hour
	SettingsRetention1Day    SettingsRetentionID = 24 * SettingsRetention1Hour
	SettingsRetention1Weeks  SettingsRetentionID = 7 * SettingsRetention1Day
	SettingsRetention1Month  SettingsRetentionID = 30 * SettingsRetention1Day
	SettingsRetention3Months SettingsRetentionID = 3 * SettingsRetention1Month
	SettingsRetention6Months SettingsRetentionID = 6 * SettingsRetention1Month
	SettingsRetention1Year   SettingsRetentionID = 12 * SettingsRetention1Month
	SettingsRetention2Years  SettingsRetentionID = 2 * SettingsRetention1Year
)

// SettingsLogs represents the settings logs of a profile.
type SettingsLogs struct {
	Enabled   bool                `json:"enabled"`
	Drop      *SettingsLogsDrop   `json:"drop,omitempty"`
	Retention SettingsRetentionID `json:"retention,omitempty"`
	Location  SettingsLocationID  `json:"location,omitempty"`
}

// SettingsPerformance represents the settings performance of a profile.
type SettingsPerformance struct {
	Ecs             bool `json:"ecs"`
	CacheBoost      bool `json:"cacheBoost"`
	CnameFlattening bool `json:"cnameFlattening"`
}
