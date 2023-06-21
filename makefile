build:
	@go build -o bin/Book_management_system

run: build
	@./bin/Book_management_system

test:
	@go test -v ./...