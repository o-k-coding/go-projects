package db

import (
	"database/sql"
	"log"

	"github.com/okeefem2/simple_bank/config"
)

const (
	dbDriver = "postgres"
)

func ConnectPostgres(config *config.Config) *sql.DB {
	dbSource, err := buildPostgresDBSource(config)
	if err != nil {
		log.Fatal("error building db source string for tests", err)
	}
	db, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	return db
}
