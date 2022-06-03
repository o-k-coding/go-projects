//go:build integration

// To use the build tag, remove the word "COMMENTED" this is there because vscode flips out with the build tag addeded
// how to run:
// cd into this dir then run
// go test . --tags integration --count=1
// --count=1 ensures no tests are cached\
// https://www.ryanchapin.com/configuring-vscode-to-use-build-tags-in-golang-to-separate-integration-and-unit-test-code/
package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	up "github.com/upper/db/v4"
)

var (
	host = "localhost"
	user = "postgres"
	password = "secret"
	dbName = "celeritas_test"
	port = "5435"
	dataSource = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var dummyUser = User{
	FirstName: "Crokus",
	LastName: "Younghand",
	Email: "cutter@phoenix.com",
	Password: "password",
	Active: 1,
}

var models Models
var resource *dockertest.Resource
var pool *dockertest.Pool


// Note future me, without using build tags, this will clash with the existing TestMain for unit tests and will cause a problem.
// Not for the tests to run for either, both build tags need to be "active"
func TestMain(m *testing.M) {
	os.Setenv("DATABASE_TYPE", "postgres")
	p, err := dockertest.NewPool("")

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Only set the package level var if we make it past the error.
	pool = p

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag: "13.4",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"}, // port inside the container
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err = pool.RunWithOptions(&opts)

	if err != nil {
		if pool != nil && resource != nil {
			_ = pool.Purge(resource)
		}
		log.Fatalf("Could not start docker resource: %s", err)
	}

	var testDB *sql.DB

	// Wait until the docker container is ready
	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dataSource, host, port, user, password, dbName))

		if err != nil {
			return err
		}

		err = testDB.Ping()
		if err != nil {
			return err
		}
		log.Printf("DB connection success!")
		return nil
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to db: %s", err)
	}

	if testDB == nil {
		_ = pool.Purge(resource)
		log.Fatalf("test DB is nil!!")
	}

	// At this point, we have the DB running and we are connected!
	err = createTables(testDB)

	if err != nil {
		_ = pool.Purge(resource) // He didn't do this in the video, but I feel like it may be important? actually should this just be deferred at the top instead? probs a good idea.
		log.Fatalf("error creating tables: %s", err)
	}

	// Init the DB models
	models = New(testDB)

	// Run the tests
	code := m.Run()

	// Cleanup docker image

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s ", err)
	}

	os.Exit(code)

}

func createTables(db *sql.DB) error {
	if db == nil {
		return errors.New("DB is nil")
	}
	stmt := `
	CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

drop table if exists users cascade;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    user_active integer NOT NULL DEFAULT 0,
    email character varying(255) NOT NULL UNIQUE,
    password character varying(60) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

drop table if exists remember_tokens;

CREATE TABLE remember_tokens (
    id SERIAL PRIMARY KEY,
    user_id integer NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    remember_token character varying(100) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON remember_tokens
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

drop table if exists tokens;

CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    user_id integer NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    first_name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    token_hash bytea NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    expiry timestamp without time zone NOT NULL
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON tokens
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
	`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

/// USER TESTS

func TestUser_Table(t *testing.T) {
	table := models.Users.Table()
	if table != "users" {
		t.Error("wrong table name returned: ", table)
	}
}

func TestUser_Insert(t *testing.T) {
	id, err := models.Users.Insert(&dummyUser)
	if err != nil {
		t.Error("failed to insert user: ", err)
	}

	if id == 0 {
		t.Error("0 returned as id after user insert")
	}
}

func TestUser_Get(t *testing.T) {
	user, err := models.Users.Get(1)
	if err != nil {
		t.Error("failed to get user: ", err)
	}

	if user.ID != 1 {
		t.Error("incorrect user id from get")
	}
}

func TestUser_GetByEmail(t *testing.T) {
	user, err := models.Users.GetByEmail("cutter@phoenix.com")
	if err != nil {
		t.Error("failed to get user by email: ", err)
	}

	if user.ID != 1 {
		t.Error("incorrect user id from get by email")
	}

	if user.Email != "cutter@phoenix.com" {
		t.Error("incorrect user email from get by email")
	}
}

func TestUser_GetAll(t *testing.T) {
	users, err := models.Users.GetAll(up.Cond{})
	if err != nil {
		t.Error("failed to get all users: ", err)
	}

	if len(users) != 1 {
		t.Error("incorrect number of users returned from get all")
	}

	if users[0].ID != 1 {
		t.Error("incorrect user id returned from GetAll")
	}
}

func TestUser_Update(t *testing.T) {
	user, err := models.Users.Get(1)

	if err != nil {
		t.Error("failed to get user: ", err)
	}

	user.Email = "crokus@phoenix.com"
	err = models.Users.Update(user)

	if err != nil {
		t.Error("failed to update user: ", err)
	}

	user, err = models.Users.Get(1)

	if err != nil {
		t.Error("failed to get updated user: ", err)
	}

	if user.Email != "crokus@phoenix.com" {
		t.Error("user email not updated")
	}
}

func TestUser_PasswordMatches(t *testing.T) {
	matches, err := models.Users.PasswordMatches("crokus@phoenix.com", "password")

	if err != nil {
		t.Error("failed to check if password matches: ", err)
	}

	if !matches {
		t.Error("password match failed")
	}

		matches, err = models.Users.PasswordMatches("crokus@phoenix.com", "incorrect_password")

		if err != nil {
		t.Error("failed to check if incorrect password does not match: ", err)
	}

	if matches {
		t.Error("incorrect password matched")
	}
}

func TestUser_ResetPassword(t *testing.T) {
	err := models.Users.ResetPassword(1, "password1")

	if err != nil {
		t.Error("failed to reset password: ", err)
	}

	matches, err := models.Users.PasswordMatches("crokus@phoenix.com", "password1")

	if err != nil {
		t.Error("failed to check password after reset: ", err)
	}

	if !matches {
		t.Error("new password failed to match after reset: ", err)
	}
}

func TestUser_Delete(t *testing.T) {
	err := models.Users.Delete(1)
	if err != nil {
		t.Error("failed to delete user: ", err)
	}

	user, err := models.Users.Get(1)

	if err == nil {
		t.Error("expected error fetching deleted user: ", err)
	}

	if user != nil {
		t.Error("failed to delete user: ", err)
	}
}

/// TOKEN TESTS

func TestToken_Table(t *testing.T) {
	table := models.Tokens.Table()
	if table != "tokens" {
		t.Error("wrong table name returned for tokens: ", table)
	}
}

func TestToken_GenerateToken(t *testing.T) {
	id, err := models.Users.Insert(&dummyUser)
	if err != nil {
		t.Error("error inserting user to generate tokens for", err)
	}

	_, err = models.Tokens.GenerateToken(id, time.Hour*24*365)
	if err != nil {
		t.Error("error generating token: ", err)
	}
}

func TestToken_Insert(t *testing.T) {
	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user to insert tokens for", err)
	}
	token, err := models.Tokens.GenerateToken(user.ID, time.Hour*24*365)
	if err != nil {
		t.Error("error generating token: ", err)
	}

	err = models.Tokens.Insert(*token, *user)

	if err != nil {
		t.Error("error inserting token: ", err)
	}
}

func TestToken_LoadUserToken(t *testing.T) {
	user, err := models.Users.GetByEmail(dummyUser.Email)
	token, err := models.Tokens.LoadUserToken(user.ID)
	if err != nil {
		t.Error("error loading token for user")
	}

	if token.UserId != user.ID || token.Email != user.Email {
		t.Error("incorrect token retrieved for user: ", token.UserId)
	}
}
