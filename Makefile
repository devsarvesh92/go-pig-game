.PHONY: test test-verbose test-coverage

# Run all tests
test:
	go test ./...

# Run tests with verbose output
test-verbose:
	go test -v ./...

# Run tests with coverage report
test-coverage:
	go test -cover ./...
	
# Run tests and generate HTML coverage report
test-coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

build:
	go build -o pig cmd/main.go