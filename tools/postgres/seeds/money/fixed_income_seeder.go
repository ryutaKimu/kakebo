package money

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func FixedIncomeSeeder(ctx context.Context, db *sql.DB, userID int64) error {
	// 固定収入：本業給与（冪等）
	_, err := db.ExecContext(ctx, `
		INSERT INTO fixed_incomes (user_id, name, amount, payment_date, memo, created_at)
		VALUES ($1, $2, $3, $4, $5, $6);
	`, userID, "本業給与", 300000, time.Date(2025, 11, 25, 0, 0, 0, 0, time.Local), "毎月の給与", time.Now())
	if err != nil {
		log.Fatalf("failed to insert fixed income: %v", err)
	}
	fmt.Println("✅ Inserted fixed income")
	if err != nil {
		return fmt.Errorf("failed to upsert fixed income: %w", err)
	}

	fmt.Println("✅ Upserted fixed income")
	return nil
}
