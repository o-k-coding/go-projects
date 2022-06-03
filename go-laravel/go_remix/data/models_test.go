package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNew(t *testing.T) {
	fakeDB, _, _ := sqlmock.New()

	defer fakeDB.Close()

	_ = os.Setenv("DATABASE_TYPE", "postgres")
	models := New(fakeDB)

	// This seems like a useless test to me, like what is the f ing point of having a compiler if you have to test this?
	if fmt.Sprintf("%T", models) != "data.Models" {
		t.Error("Wrong type returned creating new postgres pool...")
	}

	_ = os.Setenv("DATABASE_TYPE", "mysql")
	models = New(fakeDB)

	// This seems like a useless test to me, like what is the f ing point of having a compiler if you have to test this?
	if fmt.Sprintf("%T", models) != "data.Models" {
		t.Error("Wrong type returned creating new mysql pool...")
	}
}

func TestGetInsertID(t *testing.T) {
	id := int64(1) // postgres id

	returnedID := getInsertID(id)

	if returnedID != int(1) {
		t.Error("Wrong type returned for insert ID conversion")
	}

	id = 1 // mysql id

	returnedID = getInsertID(id)

	if returnedID != int(1) {
		t.Error("Wrong type returned for insert ID conversion")
	}

}
