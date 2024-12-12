build:
	@go build -o bin/main cmd/main.go

run:
	@go run cmd/main.go

clean:
	@rm -rf bin

watch:
	@air

test:
	@go test -v ./...

.PHONY: build run clean watch test
