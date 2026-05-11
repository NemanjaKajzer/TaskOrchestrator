package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file if present (local development).
	// In production, Docker Compose injects env vars directly — no .env file
	// exists in the container, so godotenv silently does nothing.
	_ = godotenv.Load()

	// From this point on, os.Getenv works identically in both environments.
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("connected to database successfully")
}
