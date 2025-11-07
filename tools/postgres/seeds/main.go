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
	dbURL := os.Getenv("GOOSE_DBSTRING")
	if dbURL == "" {
		log.Fatal("GOOSE_DBSTRING is not set")
	}

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

	var userID int64
	err = db.QueryRowContext(ctx, `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`, "テストユーザー", "test@example.com", string(hashedPassword), time.Now()).Scan(&userID)
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}
	fmt.Printf("✅ Created user (id=%d)\n", userID)

	// --- 固定収入 ---
	_, err = db.ExecContext(ctx, `
		INSERT INTO fixed_incomes (user_id, name, amount, payment_month, memo, created_at)
		VALUES ($1, $2, $3, $4, $5, $6);
	`, userID, "本業給与", 300000, "11", "毎月の給与", time.Now())
	if err != nil {
		log.Fatalf("failed to insert fixed income: %v", err)
	}
	fmt.Println("✅ Inserted fixed income")

	// --- 固定費 ---
	fixedCosts := []struct {
		Name         string
		Amount       float64
		PaymentMonth string
		Memo         string
	}{
		{"家賃", 80000, "11", "月末払い"},
		{"光熱費", 12000, "11", "電気・ガス・水道"},
		{"通信費", 8000, "11", "スマホ・Wi-Fi"},
	}

	for _, c := range fixedCosts {
		_, err := db.ExecContext(ctx, `
			INSERT INTO fixed_costs (user_id, name, amount, payment_month, memo, created_at)
			VALUES ($1, $2, $3, $4, $5, $6);
		`, userID, c.Name, c.Amount, c.PaymentMonth, c.Memo, time.Now())
		if err != nil {
			log.Fatalf("failed to insert fixed cost (%s): %v", c.Name, err)
		}
		fmt.Printf("✅ Inserted fixed cost: %s\n", c.Name)
	}

	// --- 副収入 ---
	subIncomes := []struct {
		Source       string
		Amount       float64
		PaymentMonth string
	}{
		{"Webライティング", 25000, "11"},
		{"フリマアプリ売上", 8000, "11"},
	}

	for _, si := range subIncomes {
		_, err := db.ExecContext(ctx, `
			INSERT INTO sub_incomes (user_id, name, amount, payment_month, created_at)
			VALUES ($1, $2, $3, $4, $5);
		`, userID, si.Source, si.Amount, si.PaymentMonth, time.Now())
		if err != nil {
			log.Fatalf("failed to insert sub income (%s): %v", si.Source, err)
		}
		fmt.Printf("✅ Inserted sub income: %s\n", si.Source)
	}

	// --- 収入調整 ---
	adjustments := []struct {
		Category        string
		Amount          float64
		Reason          string
		adjustmentMonth string
	}{
		{"overtime", 12000, "10月残業分", "11"},
		{"deduction", -5000, "欠勤1日", "11"},
		{"other", 3000, "交通費清算", "11"},
	}

	for _, adj := range adjustments {
		_, err := db.ExecContext(ctx, `
			INSERT INTO income_adjustments (user_id, category, amount, reason, adjustment_month, created_at)
			VALUES ($1, $2, $3, $4, $5, $6);
		`, userID, adj.Category, adj.Amount, adj.Reason, adj.adjustmentMonth, time.Now())
		if err != nil {
			log.Fatalf("failed to insert income adjustment (%s): %v", adj.Reason, err)
		}
		fmt.Printf("✅ Inserted income adjustment: %s\n", adj.Reason)
	}

	fmt.Println("🎉 Seeder finished successfully!")
}
