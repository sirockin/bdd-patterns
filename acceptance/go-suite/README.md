## go-suite

### Overview

BDD-style tests using Go's native testing framework with [testify/suite](https://github.com/stretchr/testify#suite-package), offering a fluent API that reads like Gherkin without requiring Gherkin feature files.

**Comparison with go-cucumber:**

go-cucumber uses Gherkin feature files and godog to parse and execute them:
```gherkin
Scenario: Successful sign-up
  Given Tanya has created an account
  When Tanya activates her account
  Then Tanya should be authenticated
```

go-suite expresses the same behavior in pure Go:
```go
func (s *FeatureSuite) TestSuccessfulSignUp() {
  s.
    given().personHasCreatedAnAccount("Tanya").
    when().personActivatesTheirAccount("Tanya").
    then().personShouldBeAuthenticated("Tanya")
}
```

As with the other patterns, we use a four-layer model with protocol drivers to allow reuse of the same high level specs to test different parts of the system.

### Comparison: go-cucumber vs go-suite

| Aspect | go-cucumber | go-suite |
|--------|-------------|----------|
| **Format** | Gherkin .feature files | Go test methods |
| **Runner** | godog | testify/suite |
| **Readability** | Natural language, accessible to non-programmers | Go code, readable to developers |
| **Tooling** | Requires godog, separate Gherkin parser | Standard Go tooling, IDE support |
| **Step Reuse** | Steps defined once, matched by regex | Methods can be chained fluently |
| **Refactoring** | Manual - regex patterns need updating | IDE-assisted - standard Go refactoring |
| **Discoverability** | Steps discovered at runtime via regex | Methods discovered at compile time |
| **Living Documentation** | .feature files serve as executable specs | Test code is the documentation |
| **Collaboration** | Gherkin facilitates BA/QA/Dev collaboration | Better for dev-centric teams |

### Pros and Cons

Pros:
- **Pure Go**: No additional DSL to learn, just Go
- **IDE Support**: Full refactoring, auto-completion, and navigation
- **Type Safety**: Compile-time checking of all test logic
- **Fluent API**: Chainable given/when/then/and methods provide BDD structure
- **Testify Integration**: Rich assertion library and test lifecycle hooks
- **Standard Tooling**: Works with go test, coverage, benchmarking, etc.

Cons:
- **Less Accessible**: Requires Go knowledge, not suitable for non-technical stakeholders
- **No Gherkin**: Cannot use feature files as living documentation
- **Less Separation**: Behavior and implementation are in the same language
- **Verbosity**: More code than equivalent Gherkin scenarios

### Running the Tests
From the subdirectory
```sh
cd ./acceptance/go-suite
# run tests against the http api
make test-backend
# run tests against the front end
make test-frontend
# (for go-based tests) run tests against the domain layer
make test-domain
```
### Code Organization

```
acceptance/go-suite/
├── feature_sign_up_test.go      # Sign-up feature scenarios as Go tests
├── feature_create_project_test.go # Project creation scenarios as Go tests
├── driver/                      # Protocol-specific test drivers
│   ├── driver.go               # TestDriver interface definition
│   ├── http/                   # HTTP API driver implementation
│   │   └── http.go
│   └── ui/                     # UI automation driver (Playwright)
│       └── ui.go
├── suite_test.go               # FeatureSuite setup with given/when/then fluent API
├── steps_test.go               # Step methods (reusable test building blocks)
├── main_test.go                # Test entry points (TestDomain, TestBackEnd, TestFrontEnd)
└── setup_test.go               # Server startup helpers
```

**Key Components:**

- **Feature Test Files**: Go test methods organized by feature (e.g., `feature_sign_up_test.go`)
  - Each test method represents a scenario
  - Uses fluent API for readable test structure
- **FeatureSuite**: testify/suite.Suite wrapper providing:
  - `SetupTest()` lifecycle hook to reset state before each test
  - `given()`, `when()`, `then()`, `and()` methods for fluent chaining
  - Access to testify assertions (`Assert`, `Require`)
- **Step Methods**: Reusable building blocks that:
  - Take actor names and parameters
  - Interact with the TestDriver
  - Use testify assertions for verification
  - Return the suite for method chaining
- **TestDriver Interface**: Common interface for all protocol drivers
- **Drivers**: Layer-specific implementations:
  - Domain driver (in main codebase) - tests business logic directly
  - HTTP driver - tests via REST API
  - UI driver - tests via browser automation
- **Main Tests**: Entry points that wire up the appropriate driver for each layer

### When to Use go-suite vs go-cucumber

**Choose go-suite when:**
- Your team is primarily developers comfortable with Go
- You want maximum IDE support and refactoring capabilities
- You prefer compile-time safety over runtime flexibility
- Living documentation via Gherkin is not required
- You want to minimize external dependencies

**Choose go-cucumber when:**
- You need to collaborate with non-technical stakeholders (BAs, QAs, POs)
- Gherkin feature files provide value as living documentation
- You want clear separation between behavior specification and implementation
- The team values natural language readability over code familiarity
- You're already using Cucumber in other parts of your organization
