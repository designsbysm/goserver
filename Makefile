.PHONY: all

default: test lint

benchmark:
	go test ./... -bench=. -benchmem

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html="coverage.out"

docs:
	godoc2md github.com/designsbysm/server-go | sed -e s#src/target/##g > DOCUMENTATION.md

lint:
	golangci-lint run

test:
	go test ./... -cover -race
