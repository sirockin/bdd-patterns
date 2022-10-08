## cucumber-screenplay-go

An attempt to port the officicial [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code) to `go`, using the official [godog](https://github.com/cucumber/godog/) library.

The code now replicates that of the original javascript project as such, like the javascript project and completes the use of Actor objects to implement each step. Like the original code it:
- Uses Actors with Abilities, and Actions which can be grouped to represent Tasks

Unlike the javascript project, it also uses Questions and associated helper methods

There are some differences in structure:
- `godog` does not support custom parameters so the actors are created and accessed by an `Actor(name string)` method on the `accountFeature` object
- `godog` does not support cucumber parameter syntax so regular expressions are used for these
- `go` does not support arrow functions so the implementation of actions tasks etc uses standard functions
- the implementation code is placed in a `domain` folder and package, and accessed via an `Application` interface as a first step to providing a driver-based solution
- We now inject the application into the test suite via the go TestFeatures() function so we no longer have an exported InitializeScenarios function. This means the tests can no longer be run from `godocg run` but instead should be run from `go test`

To do:
- Try providing an http/gprc implementation then separating application/domain to allow the same scenarios and specs to be used for the domain and implementation as per [go-specs-greet](https://github.com/quii/go-specs-greet)