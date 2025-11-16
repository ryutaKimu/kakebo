package model

import (
	"database/sql"
	"time"
)

type Want struct {
	ID           int64        `db:"id" json:"id"`
	UserId       int64        `db:"user_id" json:"user_id"`
	Name         string       `db:"name" json:"name"`
	TargetAmount float64      `db:"target_amount" json:"target_amount"`
	TargetDate   time.Time    `db:"target_date" json:"target_date"`
	Purchased    bool         `db:"purchased" json:"purchased"`
	PurchasedAt  sql.NullTime `db:"purchased_at" json:"purchased_at"`
	CreatedAt    time.Time    `db:"created_at" json:"created_at"`
}
