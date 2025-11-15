package money

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func SubIncomeSeeder(ctx context.Context, db *sql.DB, userID int64) error {
	subIncomes := []struct {
		Name        string
		Amount      float64
		PaymentDate time.Time
	}{
		{"Webライティング", 25000, time.Date(2025, 11, 25, 0, 0, 0, 0, time.Local)},
		{"フリマアプリ売上", 8000, time.Date(2025, 11, 25, 0, 0, 0, 0, time.Local)},
	}

	for _, si := range subIncomes {
		_, err := db.ExecContext(ctx, `
			INSERT INTO sub_incomes (user_id, name, amount, payment_date, created_at)
			VALUES ($1, $2, $3, $4, $5);
		`, userID, si.Name, si.Amount, si.PaymentDate, time.Now())
		if err != nil {
			log.Fatalf("failed to insert sub income (%s): %v", si.Name, err)
		}
		fmt.Printf("✅ Inserted sub income: %s\n", si.Name)
	}
	return nil
}
