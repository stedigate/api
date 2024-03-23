package data

import "time"

type Invoice struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	DeletedAt      time.Time `json:"-"`
	CreatedBy      int64     `json:"-"`
	UpdatedBy      int64     `json:"-"`
	ExtraID        string    `json:"extra_id,omitempty"`
	Amount         float64   `json:"amount,omitempty"`
	Currency       string    `json:"currency,omitempty"`
	WalletAddress  string    `json:"wallet_address,omitempty"`
	WalletCurrency string    `json:"wallet_currency,omitempty"`
	WalletNetwork  string    `json:"wallet_network,omitempty"`
	ConfirmedAt    time.Time `json:"confirmed_at,omitempty"`
	Type           string    `json:"type"`
	Description    string    `json:"description,omitempty,string"` // Add the string directive
	PayerEmail     []string  `json:"payer_email,omitempty"`
	PayerAmount    []string  `json:"payer_amount,omitempty"`
	PayerCurrency  []string  `json:"payer_currency,omitempty"`
	Version        int32     `json:"version"`
}
