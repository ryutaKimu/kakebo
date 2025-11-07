package model

import "time"

type FixedCost struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	Name         string    `json:"name"`
	Amount       float64   `json:"amount"`
	PaymentMonth int       `json:"payment_month"`
	Memo         string    `json:"memo,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
