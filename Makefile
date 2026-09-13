.PHONY: build run
build:
	@CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o bin/api ./cmd/api/main.go
run: build
	@./bin/api