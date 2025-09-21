.PHONY: clean build server help lint fmt vet sec test test-all test-fast coverage

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Testing targets (delegated to acceptance/Makefile)
test: ## Run tests (delegates to acceptance/Makefile)
	cd acceptance && $(MAKE) test

test-all: ## Run all tests (delegates to acceptance/Makefile)
	cd acceptance && $(MAKE) test-all

test-fast: ## Run fast tests (delegates to acceptance/Makefile)
	cd acceptance && $(MAKE) test-fast

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
	cd acceptance && go vet ./...

sec: ## Run security checks with gosec
	cd back-end && gosec ./...
	cd acceptance && gosec ./...

lint: fmt vet sec ## Run formatting and vetting

# Coverage (delegated to acceptance/Makefile)
coverage: ## Run tests with coverage (delegates to acceptance/Makefile)
	cd acceptance && $(MAKE) coverage