package services

import "github.com/andrewmzhang/nextdns/models"

type LogsQueryConfig struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Sort   string `json:"sort"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
	Device string `json:"device,omitempty"`
	Status string `json:"status,omitempty"`
	Search string `json:"search,omitempty"`
	Raw    bool   `json:"raw"`
}

type LogService[T any, C any] struct {
	path          string
	nextDNSClient *NextDNSClient
}

func (c *NextDNSClient) Logs(profileId string) *TimeSeriesService[models.Log, LogsQueryConfig] {
	return &TimeSeriesService[models.Log, LogsQueryConfig]{
		path:          "/profiles/" + profileId + "/logs",
		nextDNSClient: c,
	}
}
