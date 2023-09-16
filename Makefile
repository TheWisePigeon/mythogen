build:
	@go build .

run:
	@go run .

install-deps:
	@go mod download

clean-deps:
	@go mod tidy
