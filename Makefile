.PHONY: clean build server help lint fmt vet sec test test-all test-fast coverage

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Default subfolder if none specified
SUBFOLDER ?= cucumber

# Parameterized testing targets
test: ## Run tests (usage: make test [SUBFOLDER=cucumber|cucumber-screenplay|go-test-wrapper])
	cd acceptance/$(SUBFOLDER) && $(MAKE) test

test-all: ## Run all tests (usage: make test-all [SUBFOLDER=cucumber|cucumber-screenplay|go-test-wrapper])
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-all

test-fast: ## Run fast tests (usage: make test-fast [SUBFOLDER=cucumber|cucumber-screenplay|go-test-wrapper])
	cd acceptance/$(SUBFOLDER) && $(MAKE) test-fast

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

# Coverage target
coverage: ## Run tests with coverage (usage: make coverage [SUBFOLDER=cucumber|cucumber-screenplay|go-test-wrapper])
	cd acceptance/$(SUBFOLDER) && $(MAKE) coverage