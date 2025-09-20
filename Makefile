.PHONY: test test-domain test-http-inprocess test-http-executable test-http-docker test-fast test-integration test-all clean build server help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Individual test targets
test-domain: ## Run application unit tests (fastest)
	go test -v -run TestApplication ./acceptance

test-http-inprocess: ## Run in-process HTTP integration tests
	go test -v -run TestHTTPInProcess ./acceptance

test-http-executable: ## Run real server executable tests
	go test -v -run TestHttpExecutable ./acceptance

test-http-docker: ## Run Docker container tests (slowest)
	go test -v -run TestHttpDocker ./acceptance

# Test suites
test-fast: ## Run fast tests (application + in-process HTTP)
	go test -v -run "TestApplication|TestHTTPInProcess" ./acceptance

test-integration: ## Run all integration tests (excluding Docker)
	go test -v -run "TestHTTPInProcess|TestHttpExecutable" ./acceptance

test-all: ## Run all tests including Docker (full suite)
	go test -v ./acceptance

test: test-fast ## Default test target (fast tests only)

# Test with short mode (unit tests only)
test-short: ## Run tests in short mode (skips slow integration tests)
	go test -short -v ./acceptance

# Build targets
build: ## Build the server binary
	go build -o bin/server ./cmd/server

server: build ## Build and run the server
	./bin/server

# Clean up
clean: ## Clean build artifacts
	rm -rf bin/

# Development helpers
fmt: ## Format Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

sec: ## Run security checks with gosec
	gosec ./...

lint: fmt vet sec ## Run formatting and vetting

# Coverage
coverage: ## Run tests with coverage
	go test -coverprofile=coverage.out ./acceptance
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"