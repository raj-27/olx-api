.PHONY: build run migrate-up migrate-down

build:
	@CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o bin/api ./cmd/api/main.go
run: build
	@./bin/api

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down
