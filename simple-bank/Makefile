-include .env
export

sqlc:
	sqlc generate

migratedb:
	migrate -path db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL_MODE}" -verbose up

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/okeefem2/simple_bank/db/sqlc Store

.PHONY: server test mock
