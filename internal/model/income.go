package model

import "time"

type FixedIncome struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	Name         string    `json:"name"`
	Amount       float64   `json:"amount"`
        PaymentDate  time.Time `json:"payment_date"`
	Memo         string    `json:"memo,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
