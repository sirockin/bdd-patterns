## go-no-driver-ui

### Overview

BDD-style tests using **pure Go testing** with UI automation code directly inlined into test steps.

```go
func TestSignUp(t *testing.T) {
	ctx := setupTest(t)

	// Given
	personHasCreatedAnAccount(t, ctx, "Sue")

	// When
	personActivatesTheirAccount(t, ctx, "Sue")

	// Then
	personTriesToSignIn(t, ctx, "Sue")
	personShouldBeAuthenticated(t, ctx, "Sue")
}
```

### How It Works

Instead of abstracting browser automation behind a driver interface, this pattern:
1. Includes Playwright browser/context/page directly in the testContext struct
2. Inlines all UI automation code directly into step functions
3. Tests only run against the frontend (no domain or HTTP testing)

### Pros and Cons

Pros:
- **Absolute Simplicity**: No abstractions, no interfaces, no driver layer
- **Direct Access**: Full Playwright API available directly in steps
- **Easy to Understand**: Clear, linear code flow
- **Single Protocol**: Only tests UI, no multi-layer complexity

Cons:
- **No Abstraction**: Cannot reuse tests across different protocols (domain/HTTP/UI)
- **Frontend Only**: Cannot test business logic without a browser
- **Code Duplication**: UI automation code duplicated across all step functions
- **Tightly Coupled**: Step functions directly coupled to Playwright
- **Hard to Refactor**: Changes to UI selectors require updates in multiple places
- **No Protocol Flexibility**: Locked into browser-based testing

### When to Use This Pattern

Use this pattern when:
- You only need to test via the UI
- The test suite is small and focused
- Simplicity is more important than reusability
- You want the shortest path to UI testing

Avoid this pattern when:
- You need to test multiple protocols (domain, HTTP, UI)
- You want to run fast tests without a browser
- Your test suite is large and will benefit from abstractions
- You want flexibility to change test protocols later

### Running the Tests

```sh
cd ./acceptance/go-no-driver-ui

# Run UI tests
make test-frontend

# Or simply
make test
```

### Code Organization

```
acceptance/go-no-driver-ui/
├── feature_sign_up_test.go      # Sign-up feature tests
├── feature_create_project_test.go # Project creation tests
├── steps_test.go                # Step functions with inlined UI automation
├── main_test.go                 # TestMain setup + setupTest helper
└── setup_test.go                # Server startup helpers
```

**Key Components:**

- **Feature Test Files**: Standard Go test functions organized by feature
  - Each test calls `setupTest(t)` to get a testContext
  - Test logic calls step functions with the context
  - Comments mark Given/When/Then structure (by convention)
- **testContext**: Holds Playwright browser, context, page, and test state
  - No driver interface - just the raw Playwright objects
  - Includes lastErrors map for multi-step error scenarios
  - Provides clearAll() method for test isolation
- **Step Functions**: Regular Go functions that:
  - Take `*testing.T`, `*testContext`, and parameters
  - Use `t.Helper()` for accurate error reporting
  - Use testify assertions (`assert`, `require`)
  - Contain inlined Playwright code (no driver abstraction)
- **TestMain**: Sets up infrastructure once per test run
  - Starts frontend and backend servers
  - Provides cleanup after all tests complete
- **setupTest()**: Helper function that:
  - Creates a new testContext with Playwright browser
  - Clears all data for test isolation
  - Registers cleanup handlers

### Comparison with Other Patterns

This pattern sits at the opposite end of the spectrum from patterns with driver abstractions:

| Pattern | Abstraction | Protocols | Complexity | Flexibility |
|---------|-------------|-----------|------------|-------------|
| go-no-driver-ui | None | UI only | Lowest | Lowest |
| go-test-wrapper | Driver interface | All three | Medium | High |
| go-cucumber | Driver interface | All three | Higher | High |

Choose based on your needs for simplicity vs. flexibility.
