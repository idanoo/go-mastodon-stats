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

	gms.TIMEZONE = os.Getenv("TIMEZONE")
	if gms.TIMEZONE == "" {
		log.Fatal("TIMEZONE not set")
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
	gms.MATRIX_WEBHOOK_CHANNEL = os.Getenv("MATRIX_WEBHOOK_CHANNEL")
	if gms.MATRIX_WEBHOOK_URL == "" ||
		gms.MATRIX_WEBHOOK_API_KEY == "" ||
		gms.MATRIX_WEBHOOK_CHANNEL == "" {
		log.Println("MATRIX_WEBHOOK info incompelete. Skipping")

		// Set URL empty so we can check this later on
		gms.MATRIX_WEBHOOK_URL = ""
	}

	// Load mas1odon data if exists
	gms.MASTODON_INSTANCE_URL = os.Getenv("MASTODON_INSTANCE_URL")
	gms.MASTODON_CLIENT_ID = os.Getenv("MASTODON_CLIENT_ID")
	gms.MASTODON_CLIENT_SECRET = os.Getenv("MASTODON_CLIENT_SECRET")
	gms.MASTODON_CLIENT_USERNAME = os.Getenv("MASTODON_CLIENT_USERNAME")
	gms.MASTODON_CLIENT_PASSWORD = os.Getenv("MASTODON_CLIENT_PASSWORD")
	if gms.MASTODON_INSTANCE_URL == "" ||
		gms.MASTODON_CLIENT_ID == "" ||
		gms.MASTODON_CLIENT_SECRET == "" ||
		gms.MASTODON_CLIENT_USERNAME == "" ||
		gms.MASTODON_CLIENT_PASSWORD == "" {
		log.Println("MASTODON_INSTANCE info incompelete. Skipping")

		// Set URL empty so we can check this later on
		gms.MASTODON_INSTANCE_URL = ""
	}

	// Load schemas if set
	gms.PIXELFED_DB_SCHEMA = os.Getenv("PIXELFED_DB_SCHEMA")
	gms.MATRIX_DB_SCHEMA = os.Getenv("MATRIX_DB_SCHEMA")
	gms.MASTODON_DB_SCHEMA = os.Getenv("MASTODON_DB_SCHEMA")
	gms.MOBILIZON_DB_SCHEMA = os.Getenv("MOBILIZON_DB_SCHEMA")
	gms.PEERTUBE_DB_SCHEMA = os.Getenv("PEERTUBE_DB_SCHEMA")
	gms.BOOKWYRM_DB_SCHEMA = os.Getenv("BOOKWYRM_DB_SCHEMA")

	gms.Run()
}
