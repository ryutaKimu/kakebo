package money

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func WantsSeeder(ctx context.Context, db *sql.DB, userID int64) error {
	// 購入済みの場合
	purchasedAt := sql.NullTime{
		Time:  time.Date(2025, 10, 24, 0, 0, 0, 0, time.Local),
		Valid: true,
	}
	wants := []struct {
		Name         string
		TargetAmount float64
		TargetDate   time.Time
		Purchased    bool
		PurchasedAt  sql.NullTime
	}{
		{"フルート", 120000, time.Date(2026, 5, 1, 0, 0, 0, 0, time.Local), false, sql.NullTime{Valid: false}},
		{"Mac PC M4", 120000, time.Date(2025, 10, 25, 0, 0, 0, 0, time.Local), true, purchasedAt},
	}

	for _, wt := range wants {
		_, err := db.ExecContext(ctx,
			`INSERT INTO wants (user_id, name, target_amount, target_date, purchased, purchased_at)
			VALUES ($1, $2, $3, $4, $5, $6);
		`, userID, wt.Name, wt.TargetAmount, wt.TargetDate, wt.Purchased, wt.PurchasedAt)
		if err != nil {
			log.Fatalf("failed to insert wants table: %v", err)
		}
		fmt.Printf("✅ Inserted wants table: %s\n", wt.Name)
	}

	return nil
}
