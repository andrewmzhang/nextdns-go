package models

import "time"

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	Domain    string    `json:"domain"`
	Root      string    `json:"root"`
	Tracker   string    `json:"tracker"`
	Encrypted bool      `json:"encrypted"`
	Protocol  string    `json:"protocol"`
	ClientIp  string    `json:"clientIp"`
	Client    string    `json:"client"`
	Device    struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Model   string `json:"model"`
		LocalIp string `json:"localIp"`
	} `json:"device"`
	Status  string   `json:"status"`
	Reasons []string `json:"reasons"`
}
