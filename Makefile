.PHONY: build test coverage clean run

build:
	go build -o interpreter

test:
	go test ./... -v

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

clean:
	rm -f interpreter
	rm -f coverage.out

run:
	go run main.go

.DEFAULT_GOAL := build