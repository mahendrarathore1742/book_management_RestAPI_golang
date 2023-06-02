build:
	@go build -o bin/bank-service-rest-api

run: build
	@./bin/book_management_restapi

test:
	@go test -v ./...