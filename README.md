## bdd-patterns

## Overview

Demonstration of different BDD patterns using Cucumber/Gherkin in Go, showing how the same acceptance tests can be implemented using different organizational patterns.

This monorepo contains three main components:
- **back-end**: Go service with domain logic and HTTP API
- **front-end**: React frontend (formerly web/)
- **acceptance**: BDD acceptance tests demonstrating different patterns

This repository demonstrates two approaches to organizing BDD tests:

### 1. Cucumber (Standard Pattern)
Located in `acceptance/cucumber/` - Traditional BDD implementation with step definitions directly calling the application driver.

### 2. Cucumber-Screenplay (Screenplay Pattern)
Located in `acceptance/cucumber-screenplay/` - Implementation using the Screenplay pattern with Actors, Actions, and Questions.

Both versions run identical Gherkin scenarios against different deployment models:
- Direct domain access
- HTTP API with in-process server
- HTTP API with separate server executable
- HTTP API with Docker container
- Full UI testing with frontend and API in containers

The features are based on the official [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code) ported to Go.

We use the official [Cucumber go](https://github.com/cucumber/godog/) library to translate Gherkin to tests.

## About Screenplay

The [Screenplay Pattern](https://serenity-js.org/handbook/design/screenplay-pattern/) is an evolution of the Page Object Pattern for UI test automation. It uses Actors with Abilities to perform Actions which can be grouped into Tasks. Actors can also ask Questions to verify outcomes.

While it has its origins in UI test automation, the pattern is applicable to acceptance testing in general, including API testing as demonstrated here.

It is useful where there are multiple user roles (Actors) with different capabilities (Abilities) performing different operations (Tasks) to achieve different goals (Questions).

It carries with it an extra layer of complexity and is probably best suited to projects in organizations where BDD is already well understood and widely adopted.

## Run Tests

### Using Makefile (Recommended)

```sh
# Show all available commands
make help

# Run tests for cucumber version (default)
make test                     # Fast tests (domain + in-process HTTP)
make test-fast                # Fast tests only
make test-all                 # Full test suite including Docker

# Run tests for cucumber-screenplay version
make test-screenplay          # Fast tests (screenplay version)
make test-fast-screenplay     # Fast tests only (screenplay version)
make test-all-screenplay      # Full test suite including Docker (screenplay version)

# Run tests for both versions
make test-both                # Fast tests for both versions
make test-all-both           # Full test suite for both versions

# Coverage
make coverage                 # Coverage for cucumber version
make coverage-screenplay      # Coverage for screenplay version
```

### Direct Go Commands

```sh
# Run cucumber version tests
cd acceptance/cucumber && go test -v .

# Run cucumber-screenplay version tests
cd acceptance/cucumber-screenplay && go test -v .

# Run specific test types (from either subdirectory)
go test -v -run TestApplication .
go test -v -run TestHTTPInProcess .
go test -v -run TestHttpExecutable .
go test -v -run TestHttpDocker .
go test -v -run TestUI .
```


## Test Details

### Cucumber Version (Standard BDD)
The `acceptance/cucumber/` implementation follows standard BDD practices:
- Step definitions directly call the test driver
- Error handling is managed explicitly in step implementations
- Simple, straightforward approach suitable for most projects

### Cucumber-Screenplay Version
The `acceptance/cucumber-screenplay/` implementation demonstrates the Screenplay pattern:
- Uses Actors with Abilities, and Actions which can be grouped to represent Tasks
- Uses Questions and associated helper methods for assertions
- All scenario steps are delegated to Actor methods
- More complex but provides better abstraction for large test suites

Implementation differences from the original JavaScript project:
- `godog` does not support [cucumber expressions](https://github.com/cucumber/cucumber-expressions#readme) so:
   - regular expressions are used to map parameters as per godog examples
   - actors are created and accessed by an `Actor(name string)` method on the `suite` object
- `go` does not support arrow functions so the implementation of actions, tasks etc uses standard functions

Common architecture (both versions):
- Domain implementation code is in `back-end/internal/domain` package following Go conventions
- Public domain interfaces are exposed via `back-end/pkg/domain/` for use by acceptance tests
- HTTP server implementation is in `back-end/internal/http` package
- Test drivers in `acceptance/*/driver` provide different ways to access the domain (direct, HTTP client, UI automation)
- Application is injected into test suites via go test functions rather than exported InitializeScenarios function
- Tests run via `go test` rather than `godog run`
- Each acceptance test version has its own Go module and split into several files

## Architecture

The project follows clean architecture principles with a monorepo structure:

```
back-end/           # Go service (independent module)
├── cmd/server/     # Runnable HTTP server
├── internal/       # Internal implementation packages
│   ├── domain/     # Core business logic
│   └── http/       # HTTP server implementation
└── pkg/            # Public packages for external use
    ├── domain/     # Domain entities and services
    └── http/       # HTTP server interface

front-end/          # React frontend (formerly web/)
├── src/            # React source code
├── public/         # Static assets
└── Dockerfile      # Frontend container

acceptance/         # BDD tests
├── cucumber/       # Standard BDD implementation (independent module)
│   ├── driver/     # Test drivers for different deployment modes
│   ├── features/   # Gherkin feature files
│   └── *.go        # Step definitions and test implementation
└── cucumber-screenplay/ # Screenplay pattern implementation (independent module)
    ├── driver/     # Test drivers for different deployment modes
    ├── screenplay/ # Screenplay pattern implementation
    ├── features/   # Gherkin feature files
    └── *.go        # Step definitions, actions, questions, and test implementation
```

## Test Levels

- **Application Tests** (`make test-domain`): Direct testing of business logic (fastest ~2-3ms)
- **HTTP In-Process** (`make test-http-inprocess`): HTTP API testing with in-process server (~4-5ms)
- **Server Executable** (`make test-http-executable`): Full integration with separate server process (~1-2s)
- **Docker Container** (`make test-http-docker`): Production-like containerized testing (~30-60s)
- **UI Tests** (`make test-ui`): Full stack testing with frontend and API containers using browser automation (~60-120s)

All tests run identical BDD scenarios ensuring contract compliance across all deployment models.

### Development Workflow

```sh
# Fast feedback during development
make test-fast

# Before committing changes
make test-integration

# Full validation (CI/CD)
make test-all
```

## Build and Run Server

```sh
# Build server binary
make build

# Build and run server
make server

# Or run directly from back-end directory
cd back-end && go run ./cmd/server

# Run frontend development server
cd front-end && npm start
```

## Module Structure

Each component is now its own Go module for better dependency management:

- `back-end/go.mod` - Backend service module
- `acceptance/cucumber/go.mod` - Standard BDD acceptance tests module
- `acceptance/cucumber-screenplay/go.mod` - Screenplay pattern acceptance tests module
- `front-end/package.json` - Frontend dependencies

Both acceptance test modules import the backend as a dependency and use the backend's public API via `back-end/pkg/` packages.
