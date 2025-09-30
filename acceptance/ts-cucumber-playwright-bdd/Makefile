.PHONY: test test-http test-http-executable test-http-docker test-ui test-fast test-integration test-all test-short coverage help install build clean generate-client

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Setup targets
install: ## Install dependencies
	npm install

build: ## Build TypeScript code
	npm run build

clean: ## Clean generated files and dependencies
	npm run clean

generate-client: ## Generate TypeScript client from OpenAPI spec
	npm run generate:client

# Individual test targets
test-http-executable: ## Run HTTP tests (builds and starts backend server)
	npm run test:http-executable

test-http-docker: ## Run Docker container tests (slowest)
	npm run test:http-docker

test-ui: ## Run UI tests with frontend and API containers (not fully implemented)
	npm run test:ui

# Test suites
test-fast: ## Run fast tests (HTTP with server lifecycle management)
	npm run test:http-executable

test-integration: ## Run all integration tests (excluding Docker and UI)
	npm run test:integration

test-all: ## Run all tests including Docker and UI (full suite)
	npm run test:all

test: test-fast ## Default test target (fast tests only)

# Test with short mode (fastest tests only)
test-short: ## Run tests in short mode (HTTP tests only)
	npm run test:http-executable

# Coverage
coverage: ## Run tests with coverage
	npm run coverage

# Playwright specific targets
playwright-install: ## Install Playwright browsers
	npx playwright install

playwright-ui: ## Run Playwright tests with UI mode (HTTP tests)
	npx playwright test --ui --config=playwright-http-executable.config.ts

playwright-ui-frontend: ## Run Playwright tests with UI mode (UI tests)
	npx playwright test --ui --config=playwright-ui.config.ts

playwright-report: ## Show Playwright HTML report
	npx playwright show-report

# Development helpers
dev-setup: install generate-client playwright-install ## Complete development setup

lint: ## Run linter (if configured)
	@echo "Linting not configured yet"

typecheck: ## Run TypeScript type checking
	npx tsc --noEmit