.PHONY: test test-domain test-http-inprocess test-http-executable test-http-docker test-ui test-fast test-integration test-all test-short coverage help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Individual test targets
test-domain: ## Run application unit tests (fastest)
	RUN_APPLICATION=true RUN_BACK_END=false RUN_FRONTEND=false go test -v .

test-backend: ## Run in-process HTTP integration tests
	RUN_APPLICATION=false RUN_BACK_END=true RUN_FRONTEND=false go test -v .

test-frontend: ## Run UI tests with frontend and API containers (requires Docker)
	RUN_APPLICATION=false RUN_BACK_END=false RUN_FRONTEND=true go test -v .

# Test suites
test-all: ## Run fast tests (application + in-process HTTP)
	RUN_APPLICATION=true RUN_BACK_END=true RUN_FRONTEND=true go test -v .

test: test-domain ## Default test target (application tests only)

# Coverage
coverage: ## Run tests with coverage
	go test -coverprofile=coverage.out .
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: acceptance/go-test-wrapper/coverage.html"