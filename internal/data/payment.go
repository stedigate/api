package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/stedigate/core/pkg/postgresql"
	"time"
)

type Payment struct {
	ID               int64         `json:"id"`
	CreatedAt        time.Time     `json:"-"`
	UpdatedAt        time.Time     `json:"-"`
	DeletedAt        sql.NullTime  `json:"-"`
	ConfirmedAt      sql.NullTime  `json:"confirmed_at,omitempty"`
	ExpiredAt        time.Time     `json:"expired_at,omitempty"`
	UserID           int64         `json:"-"`
	PayoutID         sql.NullInt64 `json:"-"`
	ExtraID          string        `json:"extra_id,omitempty"`
	Type             string        `json:"type"`
	Status           string        `json:"status"`
	WalletAddress    string        `json:"wallet_address,omitempty"`
	WalletCurrency   string        `json:"wallet_currency,omitempty"`
	WalletNetwork    string        `json:"wallet_network,omitempty"`
	Amount           float64       `json:"amount,omitempty"`
	Currency         string        `json:"currency,omitempty"`
	PayCurrency      string        `json:"pay_currency,omitempty"`
	PayAmount        float64       `json:"pay_amount,omitempty"`
	PaidAmount       float64       `json:"paid_amount,omitempty"`
	CommissionAmount float64       `json:"-"`
	OutcomeAmount    float64       `json:"-"`
	OutcomeCurrency  string        `json:"-"`
	Description      string        `json:"description,omitempty"`
	Version          int32         `json:"version"`
}

// --------------------------------------------------------------------------------
// DB data

type PaymentModel struct {
	DB postgresql.DB
}

// Insert Define a method on the PaymentModel type. This method takes a single
// parameter - a pointer to a Payment struct - and returns two values. The first value
// is the ID of the newly-inserted payment, and the second value is an error.
func (m PaymentModel) Insert(payment *Payment) error {
	query := `
        INSERT INTO payments (
			amount,
			currency,
			user_id,
			extra_id,
			description,
			wallet_address,
			wallet_currency,
			wallet_network,
			pay_currency,
			pay_amount,
			expired_at,
			version
		) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 1)
        RETURNING id, created_at, version`

	args := []any{
		payment.Amount,
		payment.Currency,
		payment.UserID,
		payment.ExtraID,
		payment.Description,
		payment.WalletAddress,
		payment.WalletCurrency,
		payment.WalletNetwork,
		payment.PayCurrency,
		payment.PayAmount,
		payment.ExpiredAt,
	}
	out := []any{&payment.ID, &payment.CreatedAt, &payment.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args, out)
}

// Get Add a placeholder method for fetching a specific record from the movies table.
func (m PaymentModel) Get(id int64) (*Payment, error) {
	// Define the SQL query for retrieving the movie data.
	query := `
        SELECT 
            id, 
            created_at, 
            updated_at,
            confirmed_at, 
            expired_at, 
            user_id, 
            payout_id, 
            amount, 
            currency, 
            pay_currency, 
            pay_amount, 
            paid_amount, 
            commission_amount, 
            outcome_amount, 
            outcome_currency, 
            extra_id, 
            status, 
            wallet_address, 
            wallet_currency, 
            wallet_network, 
            description, 
            version
        FROM payments
        WHERE id = $1 and deleted_at is null`

	var payment Payment

	out := []any{
		&payment.ID,
		&payment.CreatedAt,
		&payment.UpdatedAt,
		&payment.ConfirmedAt,
		&payment.ExpiredAt,
		&payment.UserID,
		&payment.PayoutID,
		&payment.Amount,
		&payment.Currency,
		&payment.PayCurrency,
		&payment.PayAmount,
		&payment.PaidAmount,
		&payment.CommissionAmount,
		&payment.OutcomeAmount,
		&payment.OutcomeCurrency,
		&payment.ExtraID,
		&payment.Status,
		&payment.WalletAddress,
		&payment.WalletCurrency,
		&payment.WalletNetwork,
		&payment.Description,
		&payment.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, []any{id}, out)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &payment, nil
}

// Update Add a placeholder method for updating a specific record in the movies table.
func (m PaymentModel) Update(payment *Payment) error {
	return nil
}

// Delete Add a placeholder method for deleting a specific record from the movies table.
func (m PaymentModel) Delete(id int64) error {
	return nil
}

// --------------------------------------------------------------------------------
// Mocking the data

type MockPaymentModel struct{}

func (m MockPaymentModel) Insert(payment *Payment) error {
	// Mock the action...
	return nil
}

func (m MockPaymentModel) Get(id int64) (*Payment, error) {
	// Mock the action...
	return nil, nil
}

func (m MockPaymentModel) Update(payment *Payment) error {
	// Mock the action...
	return nil
}

func (m MockPaymentModel) Delete(id int64) error {
	// Mock the action...
	return nil
}
