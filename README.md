# BDD Acceptance Test Patterns

## Overview

Demonstration of a range of BDD acceptance test patterns. It is intended to:
- demonstrate the pros, cons and applicability of different BDD patterns to different use cases
- provide some usable boilerplate for quickly getting BDD acceptance tests up and running
- demonstrate using the same set of high level test specs to test different parts of the system, by injecting different protocol drivers

Currently all the examples (and the API part of the system under test) are written in `go`, but I would like to add more acceptance test patterns and more languages.

The original feature specifications are based on the official [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code).

## Build and Run Front and Back End
```sh
# Build both back end and front end
make build

# Build and run both frontend and backend concurrently
make run
```

## Run Acceptance Tests

From the subdirectory
```sh
# cd to the pattern of your choice, eg...
cd ./acceptance/go-suite

# run tests against the http api
make test-backend

# run tests against the front end
make test-frontend

# (for go-based tests) run tests against the domain layer
make test-domain

## or run help to choose a command to run a specific set of tests
make help
```

Or run the same targets from the root directory, providing the path to the SUBFOLDER as an optional parameter (default is `go-cucumber`):
```sh
# run http tests for the go-suite pattern
make test-backend SUBFOLDER=go-suite
```


## The Patterns
- **go-cucumber**: High level specs written in [Gherkin](https://cucumber.io/docs/gherkin/reference/) feature files, step definitions in Go, using [godog](https://github.com/cucumber/godog) (the official Cucumber library for Go)
- **go-cucumber-screenplay**: Same as go-cucumber but implements the [Screenplay Pattern](https://cucumber.io/docs/bdd/screenplay/) for more composable and reusable test code
- **go-suite**: Tests written in pure Go using [testify/suite](https://github.com/stretchr/testify#suite-package) with a fluent given/when/then API that reads like Gherkin
- **go-test-wrapper**: Tests written in pure Go using standard `testing` package with a wrapper function that automatically runs each test against multiple protocol layers
- **go-no-driver-api**: Tests written in pure Go with HTTP API code directly inlined into step functions (no driver abstraction layer, backend API testing only)
- **go-no-driver-ui**: Tests written in pure Go with UI automation code directly inlined into step functions (no driver abstraction layer, UI testing only)



## Common Features

The system under test (SUT) is a front end written in React that accesses a back end whose REST API is specified in `./openapi.yaml`.

The same high level specs are run for all cases and should result in identical interactions with the SUT.

### Four-Layer Model

We mostly use a [four-layer model](https://continuous-delivery.co.uk/downloads/ATDD%20Guide%2026-03-21.pdf) comprising:
1. Executable Specification: Readable specs, broken down into steps
2. Domain Specific Language: Step implementations which call...
3. Protocol Drivers: An abstraction of lowest-level interactions with the system
4. System Under Test

The protocol driver layer allows us to test different parts of the system - in our case the UI, back end http service and domain layer - using the same tests.

define a `TestDriver` interface and drivers for each layer we want to test.

The `go-no-driver-api` and `go-no-driver-ui` patterns omit the protocol driver layer and inline the HTTP/UI code directly into the step implementations. 

### Are these intermediate layers really necessary?

Abstractions and layers can help us reason about a system but too many can cause greater cognitive load.

Clearly we need 1. and 4. (otherwise we have no tests and nothing to test) but what about 2 and 3?

#### Level 2: Domain Specific Language
Whether we want to call them a Domain Specific Language or not, breaking our tests down into reusable steps with meaningful names, makes them a lot easier to reason about and easier to fix if an implementation change breaks them.

#### Level 3: Protocol Drivers

Clearly this is useful if we want to test multiple layers but what if our only product is an API or a UI, not both? For a UI you're likely to write a Page Object Model (POM) which is in itself an extra layer, albeit coupled to the page structure. For an API you could be pragmatic and use an automatically generated client. 


### Exceptions to the Four-Layer Model

- `go-no-driver-api` which omits the driver layer and inlines HTTP client code directly in step implementations. It is simpler but more brittle and can only be used to test the backend API, not the UI or domain layer.
- `go-no-driver-ui` which omits the driver layer and inlines UI automation code directly in step implementations. It is simpler but more brittle and can only be used to test the UI, not the API or domain layer.


## Contributing

Please feel free to raise issues and PRs. 

### New Example Patterns

New examples are particularly welcome. If submitting please aim for the following to enable comparison of the patterns:

- if using Gherkin, use identical feature files. Otherwise use the same feature categories, scenario wording, and if possible the same step names
- use identical interactions with the system
- where possible, as for existing examples, aim to run the same tests against UI and back end, ideally using protocol drivers. If your example is in Go, use the same protocol drivers and provide a domain test.
- provide a Makefile with the same targets where relevant. 

Thanks!


