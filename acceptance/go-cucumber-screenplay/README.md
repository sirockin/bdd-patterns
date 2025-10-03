## go-cucumber-screenplay

### Overview

BDD Specs written in Gherkin, using the official Cucumber library for go, [godog](https://github.com/cucumber/godog), combined with the **Screenplay Pattern** for more maintainable and composable test code.

The Screenplay Pattern organizes test code around:
- **Actors** - who perform actions and ask questions
- **Actions** - reusable tasks that actors can perform
- **Questions** - queries about the system state
- **Abilities** - interfaces actors use to interact with the system

As with the other patterns, we use a four-layer model with protocol drivers to allow reuse of the same high level specs to test different parts of the system.

### Screenplay Pattern Pros and Cons

Pros:
- **Composability**: Actions can be composed into higher-level actions (e.g., `signUp` combines `CreateAccount` and `Activate`)
- **Reusability**: Actions and Questions are defined once and reused across steps
- **Readability**: Step definitions become thin wrappers that read naturally
- **Actor-centric**: Multiple actors can be managed independently with their own state - especially useful for complex multi-user scenarios, especially with different types of actors
- **Expressiveness**: Fluent API (`Actor.AttemptsTo()`, `Actor.ExpectsAnswer()`) reads like natural language

Cons:
- **Learning Curve**: More concepts to understand compared to simpler patterns
- **Boilerplate**: More initial setup to define the Screenplay framework
- **Overhead**: May be overkill for very simple test scenarios


### Running the Tests
From the subdirectory
```sh
cd ./acceptance/go-cucumber-screenplay
# run tests against the http api
make test-backend
# run tests against the front end
make test-frontend
# (for go-based tests) run tests against the domain layer
make test-domain
```
### Code Organization

```
acceptance/go-cucumber-screenplay/
├── features/              # Gherkin feature files
│   ├── sign_up.feature
│   └── create_project.feature
├── screenplay/            # Screenplay pattern framework
│   └── screenplay.go     # Actor, Action, Question, Abilities types
├── driver/               # Protocol-specific test drivers
│   ├── driver.go        # TestDriver interface definition
│   ├── http/            # HTTP API driver implementation
│   │   └── http.go
│   └── ui/              # UI automation driver (Playwright)
│       └── ui.go
├── interactions_test.go  # Reusable Actions (CreateAccount, Activate, etc.)
├── questions_test.go     # Reusable Questions (amIAuthenticated, etc.)
├── suite_test.go        # Test suite setup, actor management, step registration
├── steps_test.go        # Step definitions using screenplay actions/questions
├── main_test.go         # Test entry points (TestDomain, TestBackEnd, TestFrontEnd)
└── setup_test.go        # Server startup helpers
```

**Key Components:**

- **Features**: Gherkin specs shared across all test layers
- **Screenplay Package**: Core pattern implementation
  - `Actor` - represents a user performing actions
  - `Action` - reusable behaviors (functions that take Abilities and return error)
  - `Question` - reusable queries (functions that take Abilities and return results)
  - `Abilities` - wraps the TestDriver and actor state
- **Interactions**: Library of reusable Actions composed from driver operations
- **Questions**: Library of reusable Questions for querying system state
- **Step Definitions**: Map Gherkin steps to screenplay actions/questions via actors
- **Suite**: Manages actors and registers steps with godog
- **Drivers**: Layer-specific implementations:
  - Domain driver (in main codebase) - tests business logic directly
  - HTTP driver - tests via REST API
  - UI driver - tests via browser automation
- **Main Tests**: Entry points that wire up the appropriate driver for each layer
