## cucumber-screenplay-go

## Overview

A port of the official [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code) to `go`, using the official [godog](https://github.com/cucumber/godog/) library.

## Run Tests

```sh
go test -v ./features
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

- **Domain Tests**: Direct testing of business logic (fastest)
- **HTTP In-Process**: HTTP API testing with in-process server
- **Server Executable**: Full integration with separate server process
- **Docker Container**: Production-like containerized testing

All tests run identical BDD scenarios ensuring contract compliance across all deployment models.

## To Do
- Provide a GRPC implementation and test with a new ApplicationDriver