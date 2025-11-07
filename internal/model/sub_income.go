package model

import "time"

type SubIncome struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	Source       string    `json:"source"`
	Amount       float64   `json:"amount"`
	PaymentMonth string    `json:"payment_month"`
	CreatedAt    time.Time `json:"created_at"`
}
