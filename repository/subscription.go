package repository

import (
	"database/sql"
	"dip-alert/model"
	"errors"
)

type SubscriptionRepository interface {
	Create(createSubReq model.CreateSubscriptionRequest) error
	FindAll() ([]*model.Subscription, error)
	FindBySymbol(symbol string) (*model.Subscription, error)
	Delete(symbol string) error
}

type SqliteSubscriptionRepository struct {
	db *sql.DB
}

func NewSqliteSubscriptionRepository(db *sql.DB) *SqliteSubscriptionRepository {
	return &SqliteSubscriptionRepository{db: db}
}

func (r *SqliteSubscriptionRepository) Create(createSubReq *model.CreateSubscriptionRequest) error {
	query := `
	INSERT INTO subscriptions(symbol, category, look_back_days, trigger_limit) 
		VALUES(?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, createSubReq.Symbol, createSubReq.Category, createSubReq.LookBackDays, createSubReq.TriggerLimit)

	return err
}

func (r *SqliteSubscriptionRepository) FindAll() ([]*model.Subscription, error) {
	query := `
	SELECT id, symbol, category, look_back_days, trigger_limit, last_alerted_at, created_at 
	FROM subscriptions
	ORDER BY category DESC AND symbol ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*model.Subscription
	for rows.Next() {
		subscription := &model.Subscription{}
		err := rows.Scan(
			&subscription.ID,
			&subscription.Symbol,
			&subscription.Category,
			&subscription.LookBackDays,
			&subscription.TriggerLimit,
			&subscription.LastAlertedAt,
			&subscription.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}

func (r *SqliteSubscriptionRepository) FindBySymbol(symbol string) (*model.Subscription, error)  {
	query := `
		SELECT id, symbol, category, look_back_days, trigger_limit, last_alerted_at, created_at 
		FROM subscriptions
		WHERE symbol = ?
	`

	subscription := &model.Subscription{}
	err := r.db.QueryRow(query, symbol).Scan(
		&subscription.ID,
		&subscription.Symbol,
		&subscription.Category,
		&subscription.LookBackDays,
		&subscription.TriggerLimit,
		&subscription.LastAlertedAt,
		&subscription.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("subscription not found")
	}
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (r *SqliteSubscriptionRepository) Delete(symbol string) error  {
	 _, err := r.FindBySymbol(symbol)
	if err != nil {
		return err
	}

	query := `
		DELETE FROM subscriptions
		WHERE symbol = ?
	`
	result, err := r.db.Exec(query, symbol)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowAffected == 0 {
        return errors.New("subscription not found")
    }

	return nil
}