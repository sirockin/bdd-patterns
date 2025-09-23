## bdd-patterns

## Overview

This repo demonstrates the use of a range of BDD acceptance test patterns. It is intended to:
- demonstrate the pros, const and applicability of different BDD patterns to different use cases
- provide some usable boilerplate for quickly getting BDD acceptance tests up and running
- show how we can reuse the same set of high level test specs to test different parts of the system, by injecting different protocol drivers

Currently all the examples (and the API part of the system under test) are written in `go`, but I would like to add more patterns and more languages.

The features are based on the official [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code).


## Run Tests

```sh
# cd to the pattern of your choice, eg...
cd ./acceptance/go-suite

# run all the tests
make test-all

## or run help to choose a command to run one set of tests
make help
```

## The Patterns
- go-test-wrapper: tests entirely written in `go` using test wrapper pattern to inject different protocol drivers
- go-suite: tests entirely written in `go` using suite pattern to inject different protocol drivers
- go-cucumber: high level specs written in [gherkin](https://cucumber.io/docs/gherkin/reference) steps written in go, with [godog](https://github.com/cucumber/godog/tree/main/_examples) (the official cucumber go library) used as glue
- go-cucumber-screenplay: as for go-cucumber but implements the screenplay pattern



## Common Features

The system under test (SUT) is a front end written in React that accesses a back end whose REST API is specified in `./openapi.yaml`.

The same high level specs are run for all cases and should result in identical interactions with the SUT.

We use a [four-layer model](https://continuous-delivery.co.uk/downloads/ATDD%20Guide%2026-03-21.pdf) comprising:
1. Executable Specification: Readable specs, broken down into steps
2. Domain Specific Language: Step implementations which call...
3. Protocol Drivers: An abstraction of lowest-level interactions with the system
4. System Under Test

### Protocol Drivers

The protocol driver layer allows us to test different parts of the system - in our case the UI, back end http service and domain layer - using the same tests.

To do this, in we define a `TestDriver` interface test and drivers for each layer we want to test. 


### Are these intermediate layers really necessary?

Abstractions and layers can help us reason about a system but too many can cause greater cognitive load.

Clearly we need 1. and 4. (otherwise we have no tests and nothing to test) but what about 2 and 3?

I hope to provide some examples with 2 and 3 removed (feel free to submit a PR) but for now I will argue:

- Level 2: Whether we want to call them a Domain Specific Language or not, breaking our tests down into reusable steps with meaningful names, makes them a lot easier to reason about and easier to fix if an implementation change breaks them.
- Level 3: Clearly this is useful if we want to test multiple layers but what if our only product is an API or a UI, not both? For a UI you're likely to write a Page Object Model (POM) (again to protect against brittleness). For an API you could be pragmatic and use an automatically generated client.


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


