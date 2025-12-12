default:
  @just --list

test:
  go version
  go run ./cmd/test/

format:
  go fmt ./...

lint:
  golangci-lint run ./...
