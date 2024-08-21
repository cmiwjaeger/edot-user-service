# Makefile

# Define your Go commands
build:
	@go build -o main cmd/web/main.go

run:
	@go run cmd/web/main.go

test:
	@go test ./... -v

migrate:
	@go run cmd/migrate/main.go up


migrate-down:
	@go run cmd/migrate/main.go down
