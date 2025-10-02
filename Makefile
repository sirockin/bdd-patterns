.PHONY: clean build run help lint fmt vet sec test test-all test-domain test-backend test-ui coverage

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Default subfolder if none specified
SUBFOLDER ?= go-cucumber

# Parameterized testing targets
test: ## Run tests (USAGE: make test SUBFOLDER={subfolder}, default: go-cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test

test-all: ## Run all tests (USAGE: make test-all SUBFOLDER={subfolder}, default: go-cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-all

test-domain: ## Run domain tests (USAGE: make test-domain SUBFOLDER={subfolder}, default: go-cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-domain

test-backend: ## Run backend tests with real server (USAGE: make test-backend SUBFOLDER={subfolder}, default: go-cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-backend

test-frontend: ## Run frontend tests (USAGE: make test-frontend SUBFOLDER={subfolder}, default: go-cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-frontend

test-docker: ## Run Docker tests (USAGE: make test-docker SUBFOLDER={subfolder}, default: go-cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-docker

coverage: ## Run tests with coverage (USAGE: make coverage SUBFOLDER={subfolder}, default: go-cucumber)
	cd acceptance/$(SUBFOLDER) && $(MAKE) coverage

# Build
build: build-backend build-frontend ## Build both backend and frontend

build-frontend: ## Build frontend only
	cd front-end && npm run build

build-backend: ## Build backend only
	cd back-end && make build

# Run
run: build
	@echo "Starting backend server and frontend..."
	@trap 'kill 0' EXIT; \
	cd back-end && ./bin/server & \
	cd front-end && npm start


run-frontend: build-frontend ## Build and run frontend only
	@echo "Starting frontend..."
	@cd front-end && npm start

run-backend: build-backend ## Build and run backend only
	@echo "Starting backend server..."
	@cd back-end && ./bin/server

# Clean up
clean: ## Clean build artifacts
	rm -rf back-end/bin/

# Development helpers
fmt: ## Format Go code
	go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .

vet: ## Run go vet
	cd back-end && go vet ./...
	cd acceptance/go-cucumber && go vet ./...
	cd acceptance/go-cucumber-screenplay && go vet ./...
	cd acceptance/go-test-wrapper && go vet ./...

sec: ## Run security checks with gosec
	cd back-end && gosec ./...
	cd acceptance/go-cucumber && gosec ./...
	cd acceptance/go-cucumber-screenplay && gosec ./...
	cd acceptance/go-test-wrapper && gosec ./...

lint: fmt vet sec ## Run formatting and vetting
