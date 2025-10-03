## go-test-wrapper

### Overview

BDD-style tests using **pure Go testing** with a wrapper function that automatically runs each test against multiple protocol layers (domain, HTTP API, UI). This avoids the use of a testing suite.


```go
func TestSignUp(t *testing.T) {
  withTestContext(t, func(t *testing.T, ctx *testContext) {
    // Given
    personHasCreatedAnAccount(t, ctx, "Sue")

    // When
    personActivatesTheirAccount(t, ctx, "Sue")

    // Then
    personTriesToSignIn(t, ctx, "Sue")
    personShouldBeAuthenticated(t, ctx, "Sue")
  })
}
```

This single test automatically runs as:
- `TestSignUp/Application` - tests domain logic directly
- `TestSignUp/HTTPExecutable` - tests via HTTP API
- `TestSignUp/FrontEnd` - tests via browser automation

### How It Works

The `withTestContext` wrapper function:
1. Checks `TEST_TYPE` environment variable (or runs all if unset)
2. Creates subtests for each enabled layer (Application/HTTPExecutable/FrontEnd)
3. Provides appropriate driver for each layer
4. Runs the same test logic against each driver

### Pros and Cons

Pros:
- **Minimal Dependencies**: Only uses standard `testing` package + testify assertions
- **Pure Go**: No DSL, no framework, just functions and tests
- **Maximum Simplicity**: Easiest to understand for Go developers
- **Automatic Multi-Layer Testing**: Write once, test everywhere
- **IDE Support**: Full refactoring, navigation, and debugging
- **Helper Functions**: `t.Helper()` ensures accurate failure line numbers
- **Standard Tooling**: Works seamlessly with go test, coverage, benchmarks

Cons:
- **Less Structure**: No enforced BDD format (relies on comments and discipline)
- **Use of Callback to Run Tests**: Some developers may find the testWrapper hard to reason about
- **No Fluent API**: More verbose than go-suite's method chaining
- **Environment Variables**: Layer selection via TEST_TYPE is less explicit than separate test functions
- **Two-Step Setup**: To avoid recreating the same background test environment when running multiple tests, we need to set up the server in `TestMain` and the drivers in the wrapper function.


### Running the Tests

```sh
cd ./acceptance/go-test-wrapper

# Run domain tests only (fastest)
make test-domain

# Run HTTP API tests
make test-backend

# Run UI tests
make test-frontend

# Run all tests (domain + HTTP + UI)
make test-all
```

You can also use `TEST_TYPE` environment variable directly:
```sh
TEST_TYPE=application go test -v .   # Domain layer only
TEST_TYPE=back-end go test -v .      # HTTP layer only
TEST_TYPE=front-end go test -v .     # UI layer only
go test -v .                         # All layers
```

### Code Organization

```
acceptance/go-test-wrapper/
├── feature_sign_up_test.go      # Sign-up feature tests
├── feature_create_project_test.go # Project creation tests
├── driver/                      # Protocol-specific test drivers
│   ├── driver.go               # TestDriver interface definition
│   ├── http/                   # HTTP API driver implementation
│   │   └── http.go
│   └── ui/                     # UI automation driver (Playwright)
│       └── ui.go
├── steps_test.go               # Reusable step functions + testContext
├── main_test.go                # TestMain + withTestContext wrapper
└── setup_test.go               # Server startup helpers
```

**Key Components:**

- **Feature Test Files**: Standard Go test functions organized by feature
  - Each test calls `withTestContext` wrapper
  - Test logic written once, runs against all layers
  - Comments mark Given/When/Then structure (by convention)
- **withTestContext Wrapper**: Magic function that:
  - Reads `TEST_TYPE` environment variable
  - Creates subtests for each layer
  - Injects appropriate driver via `testContext`
  - Ensures cleanup after each test
- **testContext**: Simple struct holding:
  - The `TestDriver` for the current layer
  - Shared state like `lastErrors` for multi-step scenarios
  - `clearAll()` method for test isolation
- **Step Functions**: Regular Go functions that:
  - Take `*testing.T`, `*testContext`, and parameters
  - Use `t.Helper()` for accurate error reporting
  - Use testify assertions (`assert`, `require`)
  - Are protocol-agnostic (work via TestDriver interface)
- **TestDriver Interface**: Common interface for all protocol drivers
- **Drivers**: Layer-specific implementations:
  - Domain driver (in main codebase) - tests business logic directly
  - HTTP driver - tests via REST API
  - UI driver - tests via browser automation
- **TestMain**: Sets up infrastructure once per test run
  - Starts servers based on TEST_TYPE
  - Provides cleanup after all tests complete

