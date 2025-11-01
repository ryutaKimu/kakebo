package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 環境に合わせて DB URL を変更
	dbURL := os.Getenv("GOOSE_DBSTRING")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	// --- ユーザー作成 ---
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	userQuery := `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	var userID int64
	err = db.QueryRowContext(ctx, userQuery,
		"テストユーザー",
		"test@example.com",
		string(hashedPassword),
		time.Now(),
	).Scan(&userID)
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}
	fmt.Printf("Created user with id: %d\n", userID)

	// --- 固定収入作成 ---
	incomeQuery := `
		INSERT INTO fixed_incomes (user_id, name, amount, pay_day, memo, created_at)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	_, err = db.ExecContext(ctx, incomeQuery,
		userID,
		"給料",
		300000,
		25,
		"毎月の給与",
		time.Now(),
	)
	if err != nil {
		log.Fatalf("failed to insert fixed income: %v", err)
	}
	fmt.Println("Inserted fixed income")

	// --- 固定費作成 ---
	costQuery := `
		INSERT INTO fixed_costs (user_id, name, amount, payment_date, memo, created_at)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	fixedCosts := []struct {
		Name        string
		Amount      float64
		PaymentDate int
		Memo        string
	}{
		{"家賃", 80000, 27, "月末払い"},
		{"光熱費", 12000, 15, "電気・ガス・水道"},
	}

	for _, cost := range fixedCosts {
		_, err := db.ExecContext(ctx, costQuery,
			userID,
			cost.Name,
			cost.Amount,
			cost.PaymentDate,
			cost.Memo,
			time.Now(),
		)
		if err != nil {
			log.Fatalf("failed to insert fixed cost: %v", err)
		}
		fmt.Printf("Inserted fixed cost: %s\n", cost.Name)
	}

	fmt.Println("Seeder finished successfully!")
}
