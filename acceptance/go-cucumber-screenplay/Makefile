.PHONY: test test-domain test-http test-ui test-all coverage help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Individual test targets
test-domain: ## Run application unit tests (fastest)
	go test -v -run TestApplication .

test-http: ## Run HTTP tests with real server
	go test -v -run TestHttp .

test-ui: ## Run UI tests with frontend and API containers (requires Docker)
	go test -v -run TestUI .

# Test suites
test-all: ## Run all tests including UI (full suite)
	go test -v .

test: test-domain ## Default test target (domain tests only)

# Coverage
coverage: ## Run tests with coverage
	go test -coverprofile=coverage.out .
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: acceptance/coverage.html"