## cucumber-screenplay-go

## Overview

Demonstration of the Screenplay Pattern using BDD with Cucumber/Gherkin in Go.

Also demonstrates how the same BDD scenarios can be run against different deployment models:
- Direct domain access
- HTTP API with in-process server
- HTTP API with separate server executable
- HTTP API with Docker container

Based on the official [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code) to `go`, using the official [godog](https://github.com/cucumber/godog/) library.

## Run Tests

### Using Makefile (Recommended)

```sh
# Show all available commands
make help

# Run fast tests (domain + in-process HTTP)
make test

# Run individual test types
make test-domain              # Domain unit tests (fastest)
make test-http-inprocess      # In-process HTTP integration tests
make test-http-executable     # Real server executable tests
make test-http-docker         # Docker container tests

# Run test suites
make test-fast                # Fast tests only
make test-integration         # All integration tests
make test-all                 # Full test suite including Docker
```

### Direct Go Commands

```sh
# Run all tests
go test -v ./features

# Run specific test types
go test -v -run TestDomain ./features
go test -v -run TestHTTPInProcess ./features
go test -v -run TestHttpExecutable ./features
go test -v -run TestHttpDocker ./features
```


## Test Details

The code replicates that of the original javascript project and completes the use of Actor objects to implement each step. Like the original code it:
- Uses Actors with Abilities, and Actions which can be grouped to represent Tasks

Unlike the javascript project, it also uses Questions and associated helper methods, allowing all scenario steps to be delegated to Actor methods

There are some differences in structure:
- `godog` does not seem to support [cucumber expressions](https://github.com/cucumber/cucumber-expressions#readme) so:
   - regular expressions are used to map parameters as per godog examples
   - actors are created and accessed by an `Actor(name string)` method on the `suite` object
- `go` does not support arrow functions so the implementation of actions, tasks etc uses standard functions

- to promote separation of concerns:
   - the domain implementation code is placed in the `internal/domain` package following Go conventions
   - the HTTP server implementation is in the `internal/http` package
   - test drivers in `features/driver` provide different ways to access the domain (direct, HTTP client, etc.)
   - We inject the application into the test suite via the go TestFeatures() function so we no longer have an exported InitializeScenarios function. This means the tests can no longer be run from `godog run` but instead should be run from `go test`
   - feature test code has been placed in the `features` folder and split into several files

## Architecture

The project follows clean architecture principles:

```
features/           # BDD tests and test drivers
├── driver/
│   ├── domain/     # Direct domain access driver
│   └── http/       # HTTP client driver
internal/           # Internal implementation packages
├── domain/         # Core business logic
└── http/           # HTTP server implementation
cmd/server/         # Runnable HTTP server
```

## Test Levels

- **Domain Tests** (`make test-domain`): Direct testing of business logic (fastest ~2-3ms)
- **HTTP In-Process** (`make test-http-inprocess`): HTTP API testing with in-process server (~4-5ms)
- **Server Executable** (`make test-http-executable`): Full integration with separate server process (~1-2s)
- **Docker Container** (`make test-http-docker`): Production-like containerized testing (~30-60s)

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

# Or run directly
go run ./cmd/server
```

## To Do
- Provide a GRPC implementation and test with a new ApplicationDriver