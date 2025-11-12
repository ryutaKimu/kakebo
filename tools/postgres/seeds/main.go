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
		INSERT INTO fixed_incomes (user_id, name, amount, payment_date, memo, created_at)
		VALUES ($1, $2, $3, $4, $5, $6);
	`, userID, "本業給与", 300000, time.Date(2025, 11, 25, 0, 0, 0, 0, time.Local), "毎月の給与", time.Now())
	if err != nil {
		log.Fatalf("failed to insert fixed income: %v", err)
	}
	fmt.Println("✅ Inserted fixed income")

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
			log.Fatalf("failed to insert fixed cost (%s): %v", c.Name, err)
		}
		fmt.Printf("✅ Inserted fixed cost: %s\n", c.Name)
	}

	// --- 副収入 ---
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

	// --- 収入調整 ---
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
			log.Fatalf("failed to insert income adjustment (%s): %v", adj.Reason, err)
		}
		fmt.Printf("✅ Inserted income adjustment: %s\n", adj.Reason)
	}

	// --- 目標物 ---
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
			`INSERT INTO wants (user_id, name, target_amount, target_date, purchased, purchased_at, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
		`, userID, wt.Name, wt.TargetAmount, wt.TargetDate, wt.Purchased, wt.PurchasedAt, time.Now(), time.Now())
		if err != nil {
			log.Fatalf("failed to insert wants table: %v", err)
		}
		fmt.Printf("✅ Inserted wants table: %s\n", wt.Name)
	}

	// --- 貯金 ---

	savings := []struct {
		Amount  float64
		Comment sql.NullString
	}{
		{120000, sql.NullString{String: "貯金", Valid: true}},
		{-120000, sql.NullString{String: "PC購入", Valid: true}},
	}

	for _, sav := range savings {
		_, err := db.ExecContext(ctx,
			`INSERT INTO savings (user_id, amount, comment, saved_at)
			VALUES ($1, $2, $3, $4);
		`, userID, sav.Amount, sav.Comment, time.Now())
		if err != nil {
			log.Fatalf("failed to insert savings table: %v", err)
		}
		fmt.Printf("✅ Inserted savings table: amount=%.2f\n", sav.Amount)
	}

	fmt.Println("🎉 Seeder finished successfully!")
}
