package model

import "time"

type SubIncome struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	Name         string    `json:"name"`
	Amount       float64   `json:"amount"`
	PaymentMonth int       `json:"payment_month"`
	CreatedAt    time.Time `json:"created_at"`
}
