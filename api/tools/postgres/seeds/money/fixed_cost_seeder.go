package money

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func FixedCostsSeeder(ctx context.Context, db *sql.DB, userID int64) error {
	// --- 固定費 ---
	fixedCosts := []struct {
		Name        string
		Amount      float64
		PaymentDate time.Time
		Memo        string
	}{
		{"家賃", 80000, time.Date(2025, 11, 30, 0, 0, 0, 0, time.Local), "月末払い"},
		{"光熱費", 12000, time.Date(2025, 11, 24, 0, 0, 0, 0, time.Local), "電気・ガス・水道"},
		{"通信費", 8000, time.Date(2025, 11, 26, 0, 0, 0, 0, time.Local), "スマホ・Wi-Fi"},
	}

	for _, c := range fixedCosts {
		_, err := db.ExecContext(ctx, `
			INSERT INTO fixed_costs (user_id, name, amount, payment_date, memo, created_at)
			VALUES ($1, $2, $3, $4, $5, $6);
		`, userID, c.Name, c.Amount, c.PaymentDate, c.Memo, time.Now())
		if err != nil {
			return fmt.Errorf("failed to insert fixed cost (%s): %v", c.Name, err)
		}
		fmt.Printf("✅ Inserted fixed cost: %s\n", c.Name)
	}

	return nil
}
