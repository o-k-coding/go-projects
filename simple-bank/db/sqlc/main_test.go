package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/okeefem2/simple_bank/config"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	db := ConnectPostgres(config)

	defer db.Close()

	testDB = db
	testQueries = New(testDB)

	// Can I clean up all the test sql data here?

	os.Exit(m.Run())
}
