BINARY_NAME=goRemix

build:
	@go mod vendor
	@echo "Building Go Remix..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Go Remix built!"

run: build
	@echo "Starting Go Remix..."
	@./tmp/${BINARY_NAME} &
	@echo "Go Remix started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_db:
	docker compose up -d
stop_db:
	docker compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Go Remix..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Go Remix!"

restart: stop start
