package users

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func UserSeeder(ctx context.Context, db *sql.DB) (int64, error) {
	// --- パスワードハッシュ ---
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	// --- ユーザー作成 ---
	var userID int64
	err = db.QueryRowContext(ctx, `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`, "テストユーザー", "test@example.com", string(hashedPassword), time.Now()).Scan(&userID)

	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	fmt.Printf("✅ Created user (id=%d)\n", userID)
	return userID, nil
}
