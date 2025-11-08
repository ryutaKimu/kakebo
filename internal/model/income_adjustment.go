package model

import "time"

type IncomeAdjustment struct {
	ID             int64     `db:"id" json:"id"`
	UserID         int64     `db:"user_id" json:"user_id"`
	Category       string    `db:"category" json:"category"`
	Amount         float64   `db:"amount" json:"amount"`
	Reason         string    `db:"reason" json:"reason"`
	AdjustmentDate time.Time `db:"adjustment_date" json:"adjustment_date"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
