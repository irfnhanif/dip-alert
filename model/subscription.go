package model

import "time"

type Subscription struct {
	ID            int        `json:"id"`
	Symbol        string     `json:"symbol"`
	Category      string     `json:"category"`
	LookbackDays  int        `json:"lookback_days"`
	TriggerLimit  float64    `json:"trigger_limit"`
	LastAlertedAt *time.Time `json:"last_alerted_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
}

