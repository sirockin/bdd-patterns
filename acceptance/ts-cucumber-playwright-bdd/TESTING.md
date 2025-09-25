# TypeScript Cucumber Playwright BDD Testing

This example demonstrates BDD acceptance testing using TypeScript, Cucumber features, and Playwright as the test runner using the `playwright-bdd` library.

## Overview

This implementation follows a three-layer model focused on testing the actual system:

1. **Executable Specification**: Gherkin feature files (identical to other examples)
2. **Domain Specific Language**: Step definitions implemented in TypeScript
3. **Protocol Drivers**: TypeScript implementations for HTTP and UI testing (no domain driver - tests the real backend instead)
4. **System Under Test**: The same React frontend and Go backend

## Features

- **playwright-bdd integration**: Uses the `playwright-bdd` library to run Cucumber features with Playwright
- **Protocol Drivers**: Supports testing at different layers (HTTP, UI) - tests the real backend system
- **Generated API Client**: TypeScript client auto-generated from OpenAPI specification
- **Multiple Test Environments**: Separate configurations for different test types
- **Command Line & Playwright UI**: Supports both make targets and native Playwright runner

## Setup

### Prerequisites

- Node.js 18+ and npm
- Docker (for docker and UI tests)

### Installation

```bash
# Complete development setup
make dev-setup

# Or step by step:
make install              # Install npm dependencies
make generate-client      # Generate TypeScript client from OpenAPI
make playwright-install   # Install Playwright browsers
```

## Running Tests

### Command Line (via Make targets)

```bash
# Fast feedback during development
make test-fast

# Individual test layers
make test-http-inprocess      # HTTP tests with embedded server (fastest)
make test-http-executable     # HTTP tests with external server
make test-http-docker         # HTTP tests with Docker container
make test-ui                  # UI tests with full stack

# Test suites
make test-integration         # HTTP integration tests only
make test-all                 # All tests including Docker and UI
make test-short              # Fastest HTTP tests only
```

### Playwright Runner

You can also use Playwright's native test runner:

```bash
# Run specific configuration
npx playwright test --config=playwright-http-inprocess.config.ts
npx playwright test --config=playwright-ui.config.ts

# Run all tests (uses main config)
npx playwright test

# Run with Playwright UI (interactive mode)
npx playwright test --ui

# Run specific test files
npx playwright test sign_up
npx playwright test create_project
```

### Development Workflow

```bash
# Fast feedback during development
make test-fast

# Before committing changes
make test-integration

# Full validation (CI/CD)
make test-all
```

## Test Architecture

### Protocol Drivers

The example implements two protocol drivers that test the real system:

1. **HttpDriver** (`src/drivers/http-driver.ts`): HTTP API testing using generated TypeScript client - tests the actual Go backend
2. **UIDriver** (`src/drivers/ui-driver.ts`): Full UI testing using Playwright page interactions - tests the complete React frontend + Go backend

### Step Definitions

Step definitions are organized by driver type:

- `src/steps/http-steps.ts`: HTTP-specific bindings for testing via API
- `src/steps/ui-steps.ts`: UI-specific bindings for testing via browser (future implementation)

### Configuration Files

Each test environment has its own Playwright configuration:

- `playwright-http-inprocess.config.ts`: HTTP with embedded server (fastest)
- `playwright-http-executable.config.ts`: HTTP with external server
- `playwright-http-docker.config.ts`: HTTP with Docker
- `playwright-ui.config.ts`: Full UI tests
- `playwright.config.ts`: Main configuration (all tests)

## Generated API Client

The TypeScript API client is automatically generated from the OpenAPI specification:

```bash
# Regenerate client when API changes
make generate-client

# Client code is generated in src/client/
# - models/: TypeScript interfaces for API types
# - api/: API client classes with methods for each endpoint
```

## Reports and Debugging

```bash
# Generate coverage report
make coverage

# View Playwright HTML report
make playwright-report

# Run with trace for debugging
npx playwright test --trace=on
```

## IDE Integration

### VS Code

Install the Playwright extension for VS Code to:
- Run tests directly from the editor
- Debug step definitions
- View test results inline

### Test Discovery

Playwright will automatically discover:
- Feature files in `features/` directory
- Step definitions in `src/steps/` directory
- Configuration files matching `playwright*.config.ts`

## Troubleshooting

### Common Issues

1. **"playwright-bdd deprecated" warning**: This is expected, the library is still functional
2. **Server startup issues**: Ensure the backend server can start on port 8080
3. **UI tests failing**: Check that frontend builds and starts on port 3000
4. **TypeScript errors**: Run `make typecheck` to identify type issues

### Debug Tips

```bash
# Run tests with debug output
DEBUG=playwright* npx playwright test

# Run single test with trace
npx playwright test --trace=on --grep="Successful sign-up"

# View failed test artifacts
npx playwright show-report
```

## Comparison with Other Examples

This TypeScript/Playwright example provides:

- **Same test specifications**: Uses identical Gherkin features as Go examples
- **Same architecture**: Implements the four-layer model with protocol drivers
- **Modern tooling**: Uses Playwright's advanced debugging and reporting
- **Type safety**: Full TypeScript typing from API client to test code
- **Multiple execution modes**: Both command line and Playwright UI support