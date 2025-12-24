package repository

import (
	"database/sql"
	"dip-alert/model"
)


type AlertHistoryRepository interface {
	Create(*model.CreateAlertHistoryRequest) error
}

type SqliteAlertHistoryRepository struct {
	db *sql.DB
}

func NewSqliteAlertHistoryRepository(db *sql.DB) *SqliteAlertHistoryRepository {
	return &SqliteAlertHistoryRepository{db: db}
}

func (r *SqliteAlertHistoryRepository) Create(createAlertReq *model.CreateAlertHistoryRequest) error  {
	query := `INSERT INTO alert_history(subscription_id, symbol, alert_price, peak_price, drop_percentage) 
		VALUES(?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, createAlertReq.SubscriptionID, createAlertReq.Symbol, createAlertReq.AlertPrice, createAlertReq.PeakPrice, createAlertReq.DropPercentage)

	return err
}

