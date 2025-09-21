# Testing Guide

This project implements a comprehensive testing strategy with multiple levels of testing to ensure both correctness and contract compliance. This is the **non-cucumber version** that uses pure Go tests with a fluent chaining API instead of Gherkin feature files.

## Test Architecture

Tests are written using a fluent chaining pattern that reads like natural language:

```go
s.given().
    sueHasSignedUp().
    when().
    sueCreatesAProject().
    then().
    sueShouldSeeTheProject()
```

## Test Levels

### 1. Unit Tests (Domain Logic)
**Test:** `TestApplication`
**Purpose:** Direct testing of business logic
**Speed:** Fastest (~2-3ms)

```bash
go test -v -run TestApplication
```

**Architecture:**
```
Go Test Suite → Domain Test Driver → internal/domain (in-memory)
```

### 2. Integration Tests (In-Process HTTP)
**Test:** `TestHTTPInProcess`
**Purpose:** Contract testing between HTTP API and domain
**Speed:** Fast (~4-5ms)

```bash
go test -v -run TestHTTPInProcess
```

**Architecture:**
```
Go Test Suite → HTTP Client Driver → internal/http Server (in-process) → internal/domain
```

### 3. End-to-End Tests (Real Server Process)
**Test:** `TestHttpExecutable`
**Purpose:** Full integration testing with real server executable
**Speed:** Slow (~1-2s due to process startup)

```bash
go test -v -run TestHttpExecutable
```

**Architecture:**
```
Go Test Suite → HTTP Client Driver → Server Executable (separate process) → internal/domain
```

Features:
- Builds actual server binary
- Starts server on available port
- Proper process management and cleanup using `t.Cleanup()`
- Server output monitoring for debugging
- Graceful shutdown with timeout handling

### 4. Container Tests (Docker)
**Test:** `TestHttpDocker`
**Purpose:** Production-like testing in containerized environment
**Speed:** Slowest (~30-60s due to Docker build)

```bash
# Runs automatically if Docker is available, skips if not
go test -v -run TestHttpDocker
```

**Architecture:**
```
Go Test Suite → HTTP Client Driver → Docker Container → Server Binary → internal/domain
```

Features:
- Multi-stage Docker build
- Isolated container environment
- Production-like deployment testing
- Full container lifecycle management using `t.Cleanup()`

## Running Different Test Combinations

### Development (Fast Feedback)
```bash
# Run unit and integration tests only
go test -v -run "TestApplication|TestHTTPInProcess"
```

### CI/CD Pipeline
```bash
# Run all tests (Docker test will skip if Docker not available)
go test -v
```

### Short Mode (Unit Tests Only)
```bash
go test -short -v
```

### Specific Test Types
```bash
# Unit tests only
go test -v -run TestApplication

# In-process integration tests
go test -v -run TestHTTPInProcess

# Real server integration tests
go test -v -run TestHttpExecutable

# Docker container tests (skips if Docker not available)
go test -v -run TestHttpDocker

# UI tests (full end-to-end with browser automation)
go test -v -run TestUI
```

## Test Verification Strategy

All test levels run the **identical scenarios** using Go tests with fluent chaining:
- ✅ Create one project
- ✅ Try to see someone else's project
- ✅ Successful sign-up
- ✅ Try to sign in without activating account

Example test structure:
```go
func TestCreateProjectFeature_CreateOneProject(t *testing.T) {
    driver := testhelpers.NewDomainTestDriver()
    s := NewSuite(t, driver)

    s.given().
        sueHasSignedUp().
        when().
        sueCreatesAProject().
        then().
        sueShouldSeeTheProject()
}
```

This ensures:
1. **Domain Logic Correctness** - Business rules work as expected
2. **API Contract Compliance** - HTTP API correctly implements domain behavior
3. **Server Implementation** - Actual server executable works correctly
4. **Deployment Readiness** - Containerized version functions properly

## Key Differences from Cucumber Version

- **No Gherkin files** - Tests are written directly in Go
- **Fluent chaining API** - Methods return `*suite` for method chaining
- **Direct test functions** - Each scenario is a standard Go test function
- **Better IDE support** - Full Go tooling support (autocomplete, refactoring, etc.)
- **Type safety** - Compile-time checking of test logic
- **Simpler debugging** - Standard Go debugging works seamlessly

## Test Output Summary

```
TestApplication:       ✅ 4 scenarios (2-3ms)
TestHTTPInProcess:     ✅ 4 scenarios (4-5ms)
TestHttpExecutable:    ✅ 4 scenarios (1-2s)
TestHttpDocker:        ✅ 4 scenarios (30-60s) [skipped if Docker unavailable]
TestUI:                ✅ 4 scenarios (10-30s) [skipped if Docker unavailable]
```

**Total:** 20 scenario executions across 5 different deployment models ensuring comprehensive coverage and confidence in the implementation.