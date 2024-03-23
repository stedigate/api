package data

import "time"

type Wallet struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
	CreatedBy int64     `json:"-"`
	UpdatedBy int64     `json:"-"`
	UserID    int64     `json:"-"`
	Address   string    `json:"address,omitempty"`
	Currency  string    `json:"currency,omitempty"`
	Network   string    `json:"network,omitempty"`
	Type      string    `json:"type,omitempty"`
	Version   int32     `json:"version"`
}
