.PHONY: clean build server help lint fmt vet sec test test-all test-fast test-screenplay test-all-screenplay test-fast-screenplay test-wrapper test-all-wrapper test-fast-wrapper test-both test-all-both coverage coverage-screenplay coverage-wrapper

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Testing targets for cucumber (non-screenplay version)
test: ## Run tests (cucumber version)
	cd acceptance/cucumber && $(MAKE) test

test-all: ## Run all tests (cucumber version)
	cd acceptance/cucumber && $(MAKE) test-all

test-fast: ## Run fast tests (cucumber version)
	cd acceptance/cucumber && $(MAKE) test-fast

# Testing targets for cucumber-screenplay version
test-screenplay: ## Run tests (cucumber-screenplay version)
	cd acceptance/cucumber-screenplay && $(MAKE) test

test-all-screenplay: ## Run all tests (cucumber-screenplay version)
	cd acceptance/cucumber-screenplay && $(MAKE) test-all

test-fast-screenplay: ## Run fast tests (cucumber-screenplay version)
	cd acceptance/cucumber-screenplay && $(MAKE) test-fast

# Testing targets for go-test-wrapper version
test-wrapper: ## Run tests (go-test-wrapper version)
	cd acceptance/go-test-wrapper && $(MAKE) test

test-all-wrapper: ## Run all tests (go-test-wrapper version)
	cd acceptance/go-test-wrapper && $(MAKE) test-all

test-fast-wrapper: ## Run fast tests (go-test-wrapper version)
	cd acceptance/go-test-wrapper && $(MAKE) test-fast

# Run tests for both versions
test-both: ## Run tests for both cucumber and cucumber-screenplay versions
	cd acceptance/cucumber && $(MAKE) test
	cd acceptance/cucumber-screenplay && $(MAKE) test

test-all-both: ## Run all tests for both cucumber and cucumber-screenplay versions
	cd acceptance/cucumber && $(MAKE) test-all
	cd acceptance/cucumber-screenplay && $(MAKE) test-all

# Build targets
build: ## Build the server binary
	cd back-end && go build -o bin/server ./cmd/server

server: build ## Build and run the server
	./back-end/bin/server

# Clean up
clean: ## Clean build artifacts
	rm -rf back-end/bin/

# Development helpers
fmt: ## Format Go code
	go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .

vet: ## Run go vet
	cd back-end && go vet ./...
	cd acceptance/cucumber && go vet ./...
	cd acceptance/cucumber-screenplay && go vet ./...
	cd acceptance/go-test-wrapper && go vet ./...

sec: ## Run security checks with gosec
	cd back-end && gosec ./...
	cd acceptance/cucumber && gosec ./...
	cd acceptance/cucumber-screenplay && gosec ./...
	cd acceptance/go-test-wrapper && gosec ./...

lint: fmt vet sec ## Run formatting and vetting

# Coverage targets
coverage: ## Run tests with coverage (cucumber version)
	cd acceptance/cucumber && $(MAKE) coverage

coverage-screenplay: ## Run tests with coverage (cucumber-screenplay version)
	cd acceptance/cucumber-screenplay && $(MAKE) coverage

coverage-wrapper: ## Run tests with coverage (go-test-wrapper version)
	cd acceptance/go-test-wrapper && $(MAKE) coverage