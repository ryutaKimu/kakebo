package money

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func IncomeAdjustmentsSeeder(ctx context.Context, db *sql.DB, userID int64) error {
	adjustments := []struct {
		Category       string
		Amount         float64
		Reason         string
		AdjustmentDate time.Time
	}{
		{"overtime", 12000, "10月残業分", time.Date(2025, 11, 30, 0, 0, 0, 0, time.Local)},
		{"deduction", -5000, "欠勤1日", time.Date(2025, 11, 30, 0, 0, 0, 0, time.Local)},
		{"other", 3000, "交通費清算", time.Date(2025, 11, 30, 0, 0, 0, 0, time.Local)},
	}

	for _, adj := range adjustments {
		_, err := db.ExecContext(ctx, `
			INSERT INTO income_adjustments (user_id, category, amount, reason, adjustment_date, created_at)
			VALUES ($1, $2, $3, $4, $5, $6);
		`, userID, adj.Category, adj.Amount, adj.Reason, adj.AdjustmentDate, time.Now())
		if err != nil {
			return fmt.Errorf("failed to insert income adjustment (%s): %w", adj.Reason, err)
		}
		fmt.Printf("✅ Inserted income adjustment: %s\n", adj.Reason)
	}
	return nil
}
