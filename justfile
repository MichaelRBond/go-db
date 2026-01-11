default:
  @just --list

run:
  go run ./cmd/go-db/main.go

test:
  go version
  go run ./cmd/test/

format:
  go fmt ./...

lint:
  golangci-lint run ./...
