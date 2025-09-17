## cucumber-screenplay-go

## Overview

A port of the official [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code) to `go`, using the official [godog](https://github.com/cucumber/godog/) library.

## Run Tests

```sh
go test -v ./domain
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
   - the implementation code is placed in a `domain` folder and package, and accessed via an `ApplicationDriver` interface as a first step to providing a driver-based solution
   - We now inject the application into the test suite via the go TestFeatures() function so we no longer have an exported InitializeScenarios function. This means the tests can no longer be run from `godog run` but instead should be run from `go test`
   - feature test code has been placed in the `features` folder and split into several files

## To Do
- Provide an http/grpc implementation with ApplicationDriver and testing using the same scenarios and specs to be used for the domain and implementation as per [go-specs-greet](https://github.com/quii/go-specs-greet)