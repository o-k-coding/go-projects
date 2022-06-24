package db

import (
	"database/sql"
	"log"
)

const (
	dbDriver = "postgres"
)

func ConnectPostgres() *sql.DB {
	dbSource, err := buildPostgresDBSource()
	if err != nil {
		log.Fatal("error building db source string for tests", err)
	}
	db, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	return db
}
