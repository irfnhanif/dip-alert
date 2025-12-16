package repository

import (
	"database/sql"
	"dip-alert/model"
	"log"
)

func StoreSubscription(db *sql.DB, subscription model.Subscription) {
	stmt, err := db.Prepare(`INSERT INTO subscriptions(symbol, category, look_back_days, trigger_limit, last_alerted_at) VALUES(?, ?, ?)`)
    if err != nil {
        log.Fatal(err)
    }
	defer stmt.Close()

	_, err = stmt.Exec(subscription.Symbol, subscription.Category, subscription.LookBackDays, subscription.TriggerLimit, subscription.LastAlertedAt)
	if err != nil {
        log.Fatal(err)
    }
}

func ListSubscriptions(db *sql.DB) *sql.Rows  {
	stmt, err := db.Prepare("SELECT * FROM subscriptions")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	subscriptions, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	return subscriptions
}

func DeleteSubscription(db *sql.DB, ID int)  {
	var exist int
    err := db.QueryRow("SELECT 1 FROM subscriptions WHERE id = ?", ID).Scan(&exist)
    if err != sql.ErrNoRows {
        return 
    }
    if err != nil {
        log.Fatal(err)
    }

	stmt, err := db.Prepare("DELETE FROM subscriptions WHERE id = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    stmt.Exec(ID)
}