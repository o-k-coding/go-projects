package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func buildPostgresDBSource() (string, error) {
	// In a real project I would have a package devoted to config values, and get these from that. But for now this is good enough
	err := godotenv.Load("../../.env")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	// sslMode := os.Getenv("POSTGRES_SSL_MODE")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, db), err
}
