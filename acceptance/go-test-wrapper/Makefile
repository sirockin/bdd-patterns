.PHONY: test test-domain test-http-inprocess test-http-executable test-http-docker test-ui test-fast test-integration test-all test-short coverage help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Individual test targets
test-domain: ## Run application unit tests (fastest)
	RUN_APPLICATION=true RUN_HTTP_INPROCESS=false RUN_HTTP_EXECUTABLE=false RUN_HTTP_DOCKER=false RUN_UI=false go test -v .

test-http-inprocess: ## Run in-process HTTP integration tests
	RUN_APPLICATION=false RUN_HTTP_INPROCESS=true RUN_HTTP_EXECUTABLE=false RUN_HTTP_DOCKER=false RUN_UI=false go test -v .

test-http-executable: ## Run real server executable tests
	RUN_APPLICATION=false RUN_HTTP_INPROCESS=false RUN_HTTP_EXECUTABLE=true RUN_HTTP_DOCKER=false RUN_UI=false go test -v .

test-http-docker: ## Run Docker container tests (slowest)
	RUN_APPLICATION=false RUN_HTTP_INPROCESS=false RUN_HTTP_EXECUTABLE=false RUN_HTTP_DOCKER=true RUN_UI=false go test -v .

test-ui: ## Run UI tests with frontend and API containers (requires Docker)
	RUN_APPLICATION=false RUN_HTTP_INPROCESS=false RUN_HTTP_EXECUTABLE=false RUN_HTTP_DOCKER=false RUN_UI=true go test -v .

# Test suites
test-fast: ## Run fast tests (application + in-process HTTP)
	RUN_APPLICATION=true RUN_HTTP_INPROCESS=true RUN_HTTP_EXECUTABLE=false RUN_HTTP_DOCKER=false RUN_UI=false go test -v .

test-integration: ## Run all integration tests (excluding Docker and UI)
	RUN_APPLICATION=false RUN_HTTP_INPROCESS=true RUN_HTTP_EXECUTABLE=true RUN_HTTP_DOCKER=false RUN_UI=false go test -v .

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
	@echo "Coverage report generated: acceptance/go-test-wrapper/coverage.html"