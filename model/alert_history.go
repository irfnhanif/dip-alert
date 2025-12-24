package model

import "time"

type AlertHistory struct {
	ID             int       `json:"id"`
	SubscriptionID int       `json:"subscription_id"`
	Symbol         string    `json:"symbol"`
	AlertPrice     float64   `json:"alert_price"`
	PeakPrice      float64   `json:"peak_price"`
	DropPercentage float64   `json:"drop_percentage"`
	TriggeredAt    time.Time `json:"triggered_at"`
}

type CreateAlertHistoryRequest struct {
	SubscriptionID int       `json:"subscription_id"`
	Symbol         string    `json:"symbol"`
	AlertPrice     float64   `json:"alert_price"`
	PeakPrice      float64   `json:"peak_price"`
	DropPercentage float64   `json:"drop_percentage"`
}