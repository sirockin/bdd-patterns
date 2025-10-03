## go-cucumber

### Overview

BDD Specs written in Gherkin, using the official Cucumber library for go, [godog](https://github.com/cucumber/godog) to link specs to executable code.


As with the other patterns, we use a four-layer layer model with protocol drivers to allow reuse of the same high level specs to test different parts of the system.

### Running the Tests
From the subdirectory
```sh
cd ./acceptance/go-cucumber
# run tests against the http api
make test-backend
# run tests against the front end
make test-frontend
# (for go-based tests) run tests against the domain layer
make test-domain
```
### Code Organization

```
acceptance/go-cucumber/
├── features/              # Gherkin feature files
│   ├── sign_up.feature
│   └── create_project.feature
├── driver/               # Protocol-specific test drivers
│   ├── driver.go        # TestDriver interface definition
│   ├── http/            # HTTP API driver implementation
│   │   └── http.go
│   └── ui/              # UI automation driver (Playwright)
│       └── ui.go
├── suite_test.go        # Test suite setup and step registration
├── steps_test.go        # Step definitions (Given/When/Then)
├── main_test.go         # Test entry points (TestDomain, TestBackEnd, TestFrontEnd)
└── setup_test.go        # Server startup helpers
```

**Key Components:**

- **Features**: Gherkin specs shared across all test layers
- **TestDriver Interface**: Common interface for all protocol drivers (domain, HTTP, UI)
- **Step Definitions**: Map Gherkin steps to driver actions, protocol-agnostic
- **Suite**: Registers steps and runs scenarios with godog, accepting any TestDriver
- **Drivers**: Layer-specific implementations:
  - Domain driver (in main codebase) - tests business logic directly
  - HTTP driver - tests via REST API
  - UI driver - tests via browser automation
- **Main Tests**: Entry points that wire up the appropriate driver for each layer

