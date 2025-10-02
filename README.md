# BDD Acceptance Test Patterns

## Overview

This repo demonstrates the use of a range of BDD acceptance test patterns. It is intended to:
- demonstrate the pros, cons and applicability of different BDD patterns to different use cases
- provide some usable boilerplate for quickly getting BDD acceptance tests up and running
- show how we can reuse the same set of high level test specs to test different parts of the system, by injecting different protocol drivers

Currently all the examples (and the API part of the system under test) are written in `go`, but I would like to add more patterns and more languages.

The features are based on the official [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code).

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

To do this, we define a `TestDriver` interface and drivers for each layer we want to test. 


### Are these intermediate layers really necessary?

Abstractions and layers can help us reason about a system but too many can cause greater cognitive load.

Clearly we need 1. and 4. (otherwise we have no tests and nothing to test) but what about 2 and 3?

I hope to provide some examples with 2 and 3 removed (feel free to submit a PR) but for now I will argue:

- Level 2: Whether we want to call them a Domain Specific Language or not, breaking our tests down into reusable steps with meaningful names, makes them a lot easier to reason about and easier to fix if an implementation change breaks them.
- Level 3: Clearly this is useful if we want to test multiple layers but what if our only product is an API or a UI, not both? For a UI you're likely to write a Page Object Model (POM) (again to protect against brittleness). For an API you could be pragmatic and use an automatically generated client.


## Contributing

Please feel free to raise issues and PRs. 

### New Example Patterns

New examples are particularly welcome. If submitting please aim for the following to enable comparison of the patterns:

- if using gherkin, use identical feature files. Otherwise use the same feature categories, scenario wording, and if possible the same step names
- use identical interactions with the system
- where possible, as for existing examples, aim to run the same tests against UI and back end, ideally using protocol drivers. If your example is in go, use the same protocol drivers and provide a domain test. 
- provide a makefile with the same target where relevant. 

Thanks!


