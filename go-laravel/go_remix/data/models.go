package data

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	upperDB "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper  upperDB.Session

type Models struct {
	// Any models here and in new function
	// are easily accessible in the app
	Users UserModel
	Tokens TokenModel
}

func New(pool *sql.DB) Models {
	fmt.Print("Initializating db pool")
	db = pool
	dbType := strings.ToLower(os.Getenv("DATABASE_TYPE"))
	if dbType  == "mysql" || dbType == "mariadb" {
		upper, _ = mysql.New(pool)
		fmt.Print("Using mysql pool")
	} else if dbType == "postgres" || dbType == "postgresql" {
		upper, _ = postgresql.New(pool)
		fmt.Print("Using postgres pool")
	} else {
		fmt.Print("No DB type matched, no pool created")
	}

	tokens := TokenModel{}

	return Models{
		Users: UserModel{
			tokens: tokens,
		},
		Tokens: tokens,
	}
}

// Convert the primary key ID as an int for Go models
func getInsertID(i upperDB.ID) int {
	idType := fmt.Sprintf("%T", i)
	// This handles postgres int type
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
