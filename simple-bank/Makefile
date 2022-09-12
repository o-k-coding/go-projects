-include .env
export

sqlc:
	sqlc generate

migratedb:
	migrate -path db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL_MODE}" -verbose up

migratedblatest:
	migrate -path db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL_MODE}" -verbose up 1

migratedbdown:
	migrate -path db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL_MODE}" -verbose down

migratedbdownlatest:
	migrate -path db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL_MODE}" -verbose down 1

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/okeefem2/simple_bank/db/sqlc Store

docs:
	dbdocs build doc/db.dbml

schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

# Requires
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# Will need to run go mod tidy after first run to make sure new imports are installed
generate:
	rm -f pb/*go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb \
		--grpc-gateway_opt paths=source_relative \
		--openapiv2_out doc/swagger \
    --openapiv2_opt logtostderr=true,allow_merge=true,merge_file_name=simple_bank_docs \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

.PHONY: server test mock schema docs generate
