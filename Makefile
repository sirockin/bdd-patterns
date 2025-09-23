.PHONY: clean build server help lint fmt vet sec test test-all test-fast coverage

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Default subfolder if none specified
SUBFOLDER ?= cucumber

# Parameterized testing targets
test: ## Run tests (USAGE: make test() SUBFOLDER={subfolder}), default: cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test

test-all: ## Run all tests (USAGE: make test-all (SUBFOLDER={subfolder}, default: cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-all

test-fast: ## Run fast tests (USAGE: make test-all (SUBFOLDER={subfolder}, default: cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-fast

coverage: ## Run tests with coverage (SUBFOLDER={subfolder}), default: cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) coverage

# Build targets
build: ## Build the server binary
	cd back-end && make build
	cd front-end && npm run build

run: build ## Build and run both backend server and frontend
	@echo "Starting backend server and frontend..."
	@cd back-end && ./bin/server & \
	cd front-end && npm start

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
