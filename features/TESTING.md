# Testing Guide

This project implements a comprehensive testing strategy with multiple levels of testing to ensure both correctness and contract compliance.

## Test Levels

### 1. Unit Tests (Domain Logic)
**Test:** `TestDomainFeatures`
**Purpose:** Direct testing of business logic
**Speed:** Fastest (~2-3ms)

```bash
go test -v -run TestDomainFeatures ./features
```

**Architecture:**
```
BDD Tests → Domain Driver → Domain Logic (in-memory)
```

### 2. Integration Tests (In-Process HTTP)
**Test:** `TestHTTPFeatures`
**Purpose:** Contract testing between HTTP API and domain
**Speed:** Fast (~4-5ms)

```bash
go test -v -run TestHTTPFeatures ./features
```

**Architecture:**
```
BDD Tests → HTTP Client → HTTP Server (in-process) → Domain Logic
```

### 3. End-to-End Tests (Real Server Process)
**Test:** `TestIntegrationServerExecutable`
**Purpose:** Full integration testing with real server executable
**Speed:** Slow (~1-2s due to process startup)

```bash
go test -v -run TestIntegrationServerExecutable ./features
```

**Architecture:**
```
BDD Tests → HTTP Client → Server Executable (separate process) → Domain Logic
```

Features:
- Builds actual server binary
- Starts server on available port
- Proper process management and cleanup using `t.Cleanup()`
- Server output monitoring for debugging
- Graceful shutdown with timeout handling

### 4. Container Tests (Docker)
**Test:** `TestIntegrationServerDocker`
**Purpose:** Production-like testing in containerized environment
**Speed:** Slowest (~30-60s due to Docker build)

```bash
# Runs automatically if Docker is available, skips if not
go test -v -run TestIntegrationServerDocker ./features
```

**Architecture:**
```
BDD Tests → HTTP Client → Docker Container → Server Binary → Domain Logic
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
go test -v -run "TestDomainFeatures|TestHTTPFeatures" ./features
```

### CI/CD Pipeline
```bash
# Run all tests (Docker test will skip if Docker not available)
go test -v ./features
```

### Short Mode (Unit Tests Only)
```bash
go test -short -v ./features
```

### Specific Test Types
```bash
# Unit tests only
go test -v -run TestDomainFeatures ./features

# In-process integration tests
go test -v -run TestHTTPFeatures ./features

# Real server integration tests
go test -v -run TestIntegrationServerExecutable ./features

# Docker container tests (skips if Docker not available)
go test -v -run TestIntegrationServerDocker ./features
```

## Test Verification Strategy

All test levels run the **identical BDD scenarios** (4 scenarios, 14 steps):
- ✅ Create one project
- ✅ Try to see someone else's project
- ✅ Successful sign-up
- ✅ Try to sign in without activating account

This ensures:
1. **Domain Logic Correctness** - Business rules work as expected
2. **API Contract Compliance** - HTTP API correctly implements domain behavior
3. **Server Implementation** - Actual server executable works correctly
4. **Deployment Readiness** - Containerized version functions properly

## Test Output Summary

```
TestDomainFeatures:              ✅ 4 scenarios (2-3ms)
TestHTTPFeatures:                ✅ 4 scenarios (4-5ms)
TestIntegrationServerExecutable: ✅ 4 scenarios (1-2s)
TestIntegrationServerDocker:     ✅ 4 scenarios (30-60s) [skipped if Docker unavailable]
```

**Total:** 16 scenario executions across 4 different deployment models ensuring comprehensive coverage and confidence in the implementation.