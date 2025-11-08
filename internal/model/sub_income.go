package model

import "time"

type SubIncome struct {
	ID          int64     `db:"id" json:"id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	Name        string    `db:"name" json:"name"`
	Amount      float64   `db:"amount" json:"amount"`
	PaymentDate time.Time `db:"payment_date" json:"payment_date"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
