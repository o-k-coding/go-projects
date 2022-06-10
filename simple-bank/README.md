# Simple Bank

Created following <https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/>

## Development environment setup

requires go, docker

create a .env file and add values for the following variables

```bash
DB_HOST=
POSTGRES_PASSWORD=
POSTGRES_USER=
POSTGRES_DB=
POSTGRES_PORT=
POSTGRES_SSL_MODE=
```

the docker compose file will load the .env file, so your postgres instance will be set up with the values from that file. You can then use those same values to connect.

## Postgres

### DB Diagram

Using this tool to create the DB diagram <https://dbdiagram.io/home>

specifically <https://dbdiagram.io/d/628847cdf040f104c16ba34e>

### Migrations

Using <https://github.com/golang-migrate/migrate/tree/master/cmd/migrate>

used "Go Toolchain" to install, all other methods failed for me in WSL2 lol.

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

```bash
migrate create -ext sql -dir db/migrations -seq init_schema
```

note -seq adds a sequence number to the schema.

to run migrations

- manually

```bash
source .env
migrate -path db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up
```

- with make <https://simplernerd.com/make-pass-env-nested/>

```bash
make migratedb
```

Note: might be nice to run the migrations IN docker, need a container that has go installed though etc.

in CI, would not source from .env file likely?

migration log is stored in `schema_migrations` table in the DB. If a problem occurs with a migration, I was able to resolve by removing the row for the migration and trying again after fixing the sql.
using the force flag did not work as expected 😢

### SQLC

<https://github.com/kyleconroy/sqlc>

installation

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

### SQL Driver

<https://github.com/lib/pq>

## Testing

### Deadlocks

Intuitive, but it was nice to walk through an example where he had all the queries for a transaction written out
and then using the logs to track what order the queries were running with 2 concurrent txs, opened two terminals, and two db connections, started a tx in each and executed the queries in order for each one.

<https://wiki.postgresql.org/wiki/Lock_Monitoring>

when selecting the account for update, this also locks the transfer because of the foreign key constraint

One way to fix this is to remove the fk constraints...
This is not the best solution though if you need this contraint
Instead we can tell PG that we are only selecting for NO KEY UPDATE

## DB Transactions Isolation levels

### Read Phenomena

Dirty read - tx reads data written by other uncommitted tx.

non repeatable read - tx reads same row twice and sees different value because it has been modified by another tx

phantom read - tx executes a query multiple times and gets a different set of rows each time due to another tx

serialization anomaly - result of a group of txs impossible to achieve if we run them sequentially in any order without overlapping

### 4 standard ANSI isolation levels

#### Read Uncommitted

can see data written by other uncommitted txs (dirty read, yes phantom read/non repeatable)

in PG, this behaves the same as Read Committed due to the architecture.

#### Read Committed

can only see data written by committed txns (no dirty read, yes phantom read/non repeatable)

#### Repeatable Read

same read query always returns same result (no dirty read, no phantom read, no non repeatable)

in mysql, two txs can update the same row though, and both updates will be applied.

PG throws an error if two txs try to update the same data.

#### Serializable

Can achieve the same result if txs are executes in some order rather than concurrently. (result will always be the same).

in mysql, the select will be used as a select for share, and a lock will happen on a row that is being updated. The lock has a timeout which will eventually time out. SO tx retry is needed!
also possible for a deadlock to occur if one tx selects a row that is already locked by another txn.

PG throws an error and hints for you to retry if you insert into the same table in multiple txns with the same id.
I think using a UUID from the application could be a good way to help handle this? maybe not.
