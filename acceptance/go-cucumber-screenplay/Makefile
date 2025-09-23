.PHONY: test test-domain test-http-inprocess test-http-executable test-http-docker test-ui test-fast test-integration test-all test-short coverage help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Individual test targets
test-domain: ## Run application unit tests (fastest)
	go test -v -run TestApplication .

test-http-inprocess: ## Run in-process HTTP integration tests
	go test -v -run TestHTTPInProcess .

test-http-executable: ## Run real server executable tests
	go test -v -run TestHttpExecutable .

test-http-docker: ## Run Docker container tests (slowest)
	go test -v -run TestHttpDocker .

test-ui: ## Run UI tests with frontend and API containers (requires Docker)
	go test -v -run TestUI .

# Test suites
test-fast: ## Run fast tests (application + in-process HTTP)
	go test -v -run "TestApplication|TestHTTPInProcess" .

test-integration: ## Run all integration tests (excluding Docker and UI)
	go test -v -run "TestHTTPInProcess|TestHttpExecutable" .

test-all: ## Run all tests including Docker and UI (full suite)
	go test -v .

test: test-fast ## Default test target (fast tests only)

# Test with short mode (unit tests only)
test-short: ## Run tests in short mode (skips slow integration tests)
	go test -short -v .

# Coverage
coverage: ## Run tests with coverage
	go test -coverprofile=coverage.out .
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: acceptance/coverage.html"