build:
	@go build -o bin/main cmd/main.go

run:
	@go build -o bin/main cmd/main.go
	@./bin/main

clean:
	@rm -rf bin

watch:
	@air

test:
	@go test -v ./...

.PHONY: build run clean watch test
