package entity

import (
	"time"
)

type Request struct {
	IPAddress string `json:"ip_address"`
	APIKey    string `json:"api_key"`
}

type Blocker struct {
	Request
	BlockedAt time.Time `json:"blocked_at"`
}
