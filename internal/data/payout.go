package data

import "time"

type Payout struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	DeletedAt      time.Time `json:"-"`
	CreatedBy      int64     `json:"-"`
	UpdatedBy      int64     `json:"-"`
	ExtraID        string    `json:"extra_id,omitempty"`
	Amount         float64   `json:"amount"`
	Currency       string    `json:"currency"`
	PayoutAmount   float64   `json:"payout_amount"`
	PayoutCurrency string    `json:"payout_currency"`
	WalletID       int64     `json:"wallet_id"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	ConfirmedAt    time.Time `json:"confirmed_at,omitempty"`
	Version        int32     `json:"version"`
}
