## cucumber-screenplay-go

An attempt to port the officicial [Cucumber Screenplay Example](https://github.com/cucumber-school/screenplay-example/tree/code) to `go`, using the official [godog](https://github.com/cucumber/godog/) library.

Currently, code fairly well replicates that of the original javascript project as such, like the javascript project:
- the code is only partly ported to actor objects.

There are some differences in structure:
- `godog` does not support custom parameters so the actors are created and accessed by an `Actor(name string)` method on the `accountFeature` object
- `godog` does not support cucumber parameter syntax so regular expressions are used for these
- `go` does not support arrow functions so the implementation of actions tasks etc uses standard functions
- the implementation code is placed in a `domain` folder and package, and accessed via an `Application` interface as a first step to providing a driver-based solution

To do:
- Complete the port to actor objects
- Try providing an http/gprc implementation then separating application/domain to allow the same scenarios and specs to be used for the domain and implementation as per [go-specs-greet](https://github.com/quii/go-specs-greet)