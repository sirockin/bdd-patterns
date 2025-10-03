.PHONY: test test-domain test-backend test-frontend test-all coverage help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Individual test targets
test-domain: ## Run domain tests (fastest)
	go test -v -run TestDomain .

test-backend: ## Run backend tests with real server
	go test -v -run TestBackEnd .

test-frontend: ## Run frontend tests with both backend and frontend (requires npm)
	go test -v -run TestFrontEnd .

# Test suites
test-all: ## Run all tests (full suite)
	go test -v .

test: test-domain ## Default test target (domain tests only)

# Coverage
coverage: ## Run tests with coverage
	go test -coverprofile=coverage.out .
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"