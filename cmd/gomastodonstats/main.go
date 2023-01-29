package main

import (
	gms "gomastodonstats/internal/gomastodonstats"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check Postgres info exists
	gms.POSTGRESQL_HOST = os.Getenv("POSTGRESQL_HOST")
	if gms.POSTGRESQL_HOST == "" {
		log.Fatal("POSTGRESQL_HOST not set")
	}

	gms.POSTGRESQL_USER = os.Getenv("POSTGRESQL_USER")
	if gms.POSTGRESQL_USER == "" {
		log.Fatal("POSTGRESQL_USER not set")
	}

	gms.POSTGRESQL_PASS = os.Getenv("POSTGRESQL_PASS")
	if gms.POSTGRESQL_PASS == "" {
		log.Fatal("POSTGRESQL_PASS not set")
	}

	gms.POSTGRESQL_STATS_DB = os.Getenv("POSTGRESQL_STATS_DB")
	if gms.POSTGRESQL_STATS_DB == "" {
		log.Fatal("POSTGRESQL_STATS_DB not set")

	}

	// Load matrix data if exists, else set URL blank
	gms.MATRIX_WEBHOOK_URL = os.Getenv("MATRIX_WEBHOOK_URL")
	gms.MATRIX_WEBHOOK_API_KEY = os.Getenv("MATRIX_WEBHOOK_API_KEY")
	gms.MATRIX_WEBHOOK_CHANNEL = os.Getenv("MATRIX_ACCOUNT_CHANNEL")
	if gms.MATRIX_WEBHOOK_URL == "" ||
		gms.MATRIX_WEBHOOK_API_KEY == "" ||
		gms.MATRIX_WEBHOOK_CHANNEL == "" {
		log.Println("MATRIX_WEBHOOK info incompelete. Skipping")

		// Set URL empty so we can check this later on
		gms.MATRIX_WEBHOOK_URL = ""
	}

	// Load schemas if set
	gms.PIXELFED_DB_SCHEMA = os.Getenv("PIXELFED_DB_SCHEMA")
	gms.MATRIX_DB_SCHEMA = os.Getenv("MATRIX_DB_SCHEMA")
	gms.MASTODON_DB_SCHEMA = os.Getenv("MASTODON_DB_SCHEMA")
	gms.MOBILIZON_DB_SCHEMA = os.Getenv("MOBILIZON_DB_SCHEMA")
	gms.PEERTUBE_DB_SCHEMA = os.Getenv("PEERTUBE_DB_SCHEMA")

	gms.Run()
}
