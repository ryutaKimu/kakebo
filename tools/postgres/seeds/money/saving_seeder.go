package money

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func SavingSeeder(ctx context.Context, db *sql.DB, userID int64) error {
	savings := []struct {
		Amount  float64
		Comment sql.NullString
		SavedAt time.Time
	}{
		{120000, sql.NullString{String: "貯金", Valid: true}, time.Date(2025, 11, 24, 0, 0, 0, 0, time.Local)},
		{-120000, sql.NullString{String: "PC購入", Valid: true}, time.Date(2025, 11, 24, 0, 0, 0, 0, time.Local)},
	}

	for _, sav := range savings {
		_, err := db.ExecContext(ctx,
			`INSERT INTO savings (user_id, amount, comment, saved_at)
			VALUES ($1, $2, $3, $4);
		`, userID, sav.Amount, sav.Comment, sav.SavedAt)
		if err != nil {
			return fmt.Errorf("failed to insert savings table: %w", err)
		}
		fmt.Printf("✅ Inserted savings table: amount=%.2f\n", sav.Amount)
	}
	return nil
}
