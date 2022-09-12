package db

import (
	"fmt"

	"github.com/okeefem2/simple_bank/config"
)

// Another approach would be to contain this in the env file itself lol.
// But at least this way, there are more fine grained controls?
func BuildPostgresDBSource(config *config.Config) (string, error) {
	user := config.PostgresUser
	password := config.PostgresPassword
	host := config.PostgresHost
	port := config.PostgresPort
	db := config.PostgresDB
	sslMode := config.PostgresSSLMode
	conn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, db, sslMode)
	return conn, nil
}
