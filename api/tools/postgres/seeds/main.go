package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ryutaKimu/kakebo/api/tools/postgres/seeds/dbutil"
	"github.com/ryutaKimu/kakebo/api/tools/postgres/seeds/money"
	"github.com/ryutaKimu/kakebo/api/tools/postgres/seeds/users"
)

func main() {
	db, err := dbutil.Connect()
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	userID, err := users.UserSeeder(ctx, db)
	if err != nil {
		log.Fatalf("UserSeeder failed: %v", err)
	}

	if err := money.FixedIncomeSeeder(ctx, db, userID); err != nil {
		log.Fatalf("FixedIncomeSeeder failed: %v", err)
	}

	if err := money.FixedCostsSeeder(ctx, db, userID); err != nil {
		log.Fatalf("FixedCostsSeeder failed: %v", err)
	}

	if err := money.SubIncomeSeeder(ctx, db, userID); err != nil {
		log.Fatalf("SubIncomeSeeder failed: %v", err)
	}

	if err := money.IncomeAdjustmentsSeeder(ctx, db, userID); err != nil {
		log.Fatalf("IncomeAdjustmentsSeeder failed: %v", err)
	}

	if err := money.WantsSeeder(ctx, db, userID); err != nil {
		log.Fatalf("WantsSeeder failed: %v", err)
	}

	if err := money.SavingSeeder(ctx, db, userID); err != nil {
		log.Fatalf("SavingSeeder failed: %v", err)
	}

	fmt.Println("🎉 Seeder finished successfully!")
}
