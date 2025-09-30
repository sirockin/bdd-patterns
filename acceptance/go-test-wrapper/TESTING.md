# Test Wrapper Pattern

This directory demonstrates the **Test Wrapper Pattern** - a testing approach that injects different test drivers into the same test functions without using a test suite framework.

## What is the Test Wrapper Pattern?

The Test Wrapper Pattern is a testing approach where:

1. **Test functions are written once** - Each test scenario is defined as a single function
2. **Driver injection** - Different test drivers (application, HTTP, Docker, UI) are injected via a wrapper function
3. **Environment-based control** - Environment variables control which test drivers run
4. **No test suite dependency** - Unlike the go-suite pattern, this doesn't depend on testify/suite

## Key Components

### 1. withTestContext Wrapper Function

The core of the pattern is the `withTestContext` function that:

```go
func withTestContext(t *testing.T, testFn func(t *testing.T, driver driver.TestDriver)) {
    // Check environment variables to determine which drivers to run
    runApplication := os.Getenv("RUN_APPLICATION") != "false"
    runHTTPInProcess := os.Getenv("RUN_HTTP_INPROCESS") != "false"
    // ... other driver checks

    if runApplication {
        t.Run("Application", func(t *testing.T) {
            testFn(t, testhelpers.NewDomainTestDriver())
        })
    }
    // ... other driver runs
}
```

### 2. Test Functions

Test functions are written using the wrapper:

```go
func TestCreateOneProject(t *testing.T) {
    withTestContext(t, func(t *testing.T, testDriver driver.TestDriver) {
        // Arrange
        ctx := newTestContext(testDriver)
        defer ctx.clearAll()

        // Given
        assertNoError(t, ctx.personHasSignedUp("Sue"))

        // When
        assertNoError(t, ctx.personCreatesAProject("Sue"))

        // Then
        assertNoError(t, ctx.personShouldSeeTheirProject("Sue"))
    })
}
```

### 3. Test Context

The `testContext` encapsulates:
- The test driver
- Error state management
- Step definition methods

### 4. Step Definitions

Step definitions are methods on the test context:

```go
func (ctx *testContext) personHasSignedUp(name string) error {
    if err := ctx.driver.CreateAccount(name); err != nil {
        return err
    }
    if _, err := ctx.driver.GetAccount(name); err != nil {
        return nil
    }
    return ctx.driver.Activate(name)
}
```

## Environment Variable Control

The pattern uses environment variables to control which test drivers run:

| Variable | Purpose |
|----------|---------|
| `RUN_APPLICATION` | Run tests against the domain model (fastest) |
| `RUN_HTTP_INPROCESS` | Run tests against in-process HTTP server |
| `RUN_HTTP` | Run tests against real server executable |
| `RUN_HTTP_DOCKER` | Run tests against Docker container |
| `RUN_UI` | Run tests against UI with browser automation |

## Makefile Targets

The pattern supports the same Makefile targets as other patterns:

```bash
make test-domain          # Application tests only
make test-http-inprocess  # In-process HTTP tests only
make test-http-executable # Real server executable tests
make test-http-docker     # Docker container tests
make test-ui              # UI automation tests
make test-fast            # Application + in-process HTTP
make test-integration     # HTTP tests without Docker/UI
make test-all             # All tests including Docker and UI
```

## Advantages

1. **No suite dependency** - Doesn't require testify/suite or similar frameworks
2. **Flexible control** - Environment variables provide fine-grained control
3. **Standard Go testing** - Uses standard Go testing tools and conventions
4. **Parallel execution** - Each driver runs as a subtest, enabling parallel execution
5. **Clear separation** - Test logic is separate from driver management

## Disadvantages

1. **Environment variable complexity** - Requires managing multiple environment variables
2. **Less structured** - No formal setup/teardown lifecycle compared to suites
3. **Manual context management** - Test context must be manually created and cleaned up

## Test Output

When running tests, you'll see output like:

```
=== RUN   TestCreateOneProject
=== RUN   TestCreateOneProject/Application
=== RUN   TestCreateOneProject/HTTPInProcess
--- PASS: TestCreateOneProject (0.11s)
    --- PASS: TestCreateOneProject/Application (0.00s)
    --- PASS: TestCreateOneProject/HTTPInProcess (0.11s)
```

This shows that each test function runs against multiple drivers as subtests.

## Comparison with Other Patterns

| Pattern | Suite Framework | Driver Control | Setup/Teardown |
|---------|----------------|----------------|-----------------|
| Cucumber | Gherkin+godog | Feature tags | Hooks |
| Screenplay | Gherkin+godog | Feature tags | Hooks |
| Go Suite | testify/suite | Go build tags | Suite methods |
| **Wrapper** | **None** | **Environment vars** | **Manual** |

## Best Practices

1. **Always use defer for cleanup** - `defer ctx.clearAll()` ensures cleanup
2. **Use assertNoError helper** - Provides better error messages than direct assertions
3. **Keep test functions focused** - Each test should verify a single behavior
4. **Use meaningful environment variable combinations** - The Makefile targets provide good defaults
5. **Consider parallel execution** - The wrapper pattern naturally supports parallel subtests

## Usage

To run tests with specific drivers:

```bash
# Run only application tests
RUN_APPLICATION=true RUN_HTTP_INPROCESS=false go test -v .

# Run application and in-process HTTP tests
make test-fast

# Run all tests
make test-all
```

This pattern provides a lightweight alternative to test suites while maintaining the ability to run the same tests against multiple implementations.