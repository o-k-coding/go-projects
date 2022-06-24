package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	db := ConnectPostgres()

	defer db.Close()

	testDB = db
	testQueries = New(testDB)

	// Can I clean up all the test sql data here?

	os.Exit(m.Run())
}
