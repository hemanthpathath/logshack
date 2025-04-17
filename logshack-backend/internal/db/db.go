package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPool(databaseURL string) {
	var err error
	Pool, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func Init() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("❌ DB ping failed: %v", err)
	}

	fmt.Println("✅ Connected to Postgres!")
}
