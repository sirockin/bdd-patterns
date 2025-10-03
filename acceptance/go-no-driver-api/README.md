## go-no-driver-api

### Overview

BDD-style tests using **pure Go testing** with HTTP API code directly inlined into test steps. 

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

Instead of abstracting HTTP client calls behind a driver interface, this pattern:
1. Includes HTTP client and base URL directly in the testContext struct
2. Inlines all HTTP API code directly into step functions
3. Tests only run against the backend API (no domain or UI testing)

### Pros and Cons

Pros:
- **Absolute Simplicity**: No abstractions, no interfaces, no driver layer
- **Direct Access**: Full HTTP client API available directly in steps
- **Easy to Understand**: Clear, linear code flow
- **Single Protocol**: Only tests HTTP API, no multi-layer complexity
- **Fast Execution**: No browser overhead, direct HTTP requests

Cons:
- **No Abstraction**: Cannot reuse tests across different protocols (domain/HTTP/UI)
- **Backend Only**: Cannot test UI interactions
- **Code Duplication**: HTTP client code duplicated across all step functions
- **Tightly Coupled**: Step functions directly coupled to HTTP client
- **Hard to Refactor**: Changes to API endpoints require updates in multiple places
- **No Protocol Flexibility**: Locked into HTTP-based testing

### When to Use This Pattern

Use this pattern when:
- You only need to test via HTTP API
- The test suite is small and focused
- Simplicity is more important than reusability
- You want the shortest path to API testing
- You want fast execution without browser overhead

Avoid this pattern when:
- You need to test multiple protocols (domain, HTTP, UI)
- You want to test UI interactions
- Your test suite is large and will benefit from abstractions
- You want flexibility to change test protocols later

### Running the Tests

```sh
cd ./acceptance/go-no-driver-api

# Run API tests
make test-backend

# Or simply
make test
```

### Code Organization

```
acceptance/go-no-driver-api/
├── feature_sign_up_test.go      # Sign-up feature tests
├── feature_create_project_test.go # Project creation tests
├── steps_test.go                # Step functions with inlined HTTP API code
├── main_test.go                 # TestMain setup + setupTest helper
└── setup_test.go                # Server startup helpers + testContext
```

**Key Components:**

- **Feature Test Files**: Standard Go test functions organized by feature
  - Each test calls `setupTest(t)` to get a testContext
  - Test logic calls step functions with the context
  - Comments mark Given/When/Then structure (by convention)
- **testContext**: Holds HTTP client, base URL, and test state
  - No driver interface - just the raw HTTP client
  - Includes lastErrors map for multi-step error scenarios
  - Provides clearAll() method for test isolation
- **Step Functions**: Regular Go functions that:
  - Take `*testing.T`, `*testContext`, and parameters
  - Use `t.Helper()` for accurate error reporting
  - Use testify assertions (`assert`, `require`)
  - Contain inlined HTTP client code (no driver abstraction)
- **TestMain**: Sets up infrastructure once per test run
  - Starts backend server
  - Provides cleanup after all tests complete
- **setupTest()**: Helper function that:
  - Creates a new testContext with HTTP client
  - Clears all data for test isolation
  - Registers cleanup handlers

### Comparison with Other Patterns

This pattern sits at the opposite end of the spectrum from patterns with driver abstractions:

| Pattern | Abstraction | Protocols | Complexity | Flexibility |
|---------|-------------|-----------|------------|-------------|
| go-no-driver-api | None | HTTP API only | Lowest | Lowest |
| go-no-driver-ui | None | UI only | Lowest | Lowest |
| go-test-wrapper | Driver interface | All three | Medium | High |
| go-cucumber | Driver interface | All three | Higher | High |

Choose based on your needs for simplicity vs. flexibility.
