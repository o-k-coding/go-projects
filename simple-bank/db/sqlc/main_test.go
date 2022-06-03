package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
const (
	dbDriver = "postgres"
)

func TestMain(m *testing.M) {
	dbSource, err := buildPostgresDBSource()
	if (err != nil) {
		log.Fatal("error building db source string for tests", err)
	}
	db, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	defer db.Close()

	testDB = db
	testQueries = New(testDB)

	// Can I clean up all the test sql data here?

	os.Exit(m.Run())
}
