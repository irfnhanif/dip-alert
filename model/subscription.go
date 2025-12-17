package model

import "time"

type Subscription struct {
	ID            int        `json:"id"`
	Symbol        string     `json:"symbol"`
	Category      string     `json:"category"`
	LookBackDays  int        `json:"look_back_days"`
	TriggerLimit  float64    `json:"trigger_limit"`
	LastAlertedAt *time.Time `json:"last_alerted_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
}

type CreateSubscriptionRequest struct {
	Symbol        string     `json:"symbol"`
	Category      string     `json:"category"`
	LookBackDays  int        `json:"look_back_days"`
	TriggerLimit  float64    `json:"trigger_limit"`
}

type UpdateSubscriptionRequest struct {
	Symbol        string     `json:"symbol"`
	Category      string     `json:"category"`
	LookBackDays  int        `json:"look_back_days"`
	TriggerLimit  float64    `json:"trigger_limit"`
}

