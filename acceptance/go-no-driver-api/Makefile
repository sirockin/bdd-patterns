.PHONY: test-backend test coverage help

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

test-backend: ## Run backend API tests (no driver pattern)
	go test -v .

test: test-backend ## Default test target (backend tests only)

# Coverage
coverage: ## Run tests with coverage
	go test -coverprofile=coverage.out .
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
