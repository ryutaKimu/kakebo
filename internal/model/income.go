package model

import "time"

type FixedIncome struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	PayDay    int       `json:"pay_day"`
	Memo      string    `json:"memo,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
