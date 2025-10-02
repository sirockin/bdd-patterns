package features_test

import (
	"os"
	"testing"
	"time"

	httpdriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/http"
	uidriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/ui"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"
)

var (
	serverURL   string
	frontendURL string
)

func TestMain(m *testing.M) {
	cleanup := func() {}

	switch os.Getenv("TEST_TYPE") {
	case "application":
		// No setup needed for application tests
	case "back-end":
		serverURL, cleanup = startServerExecutable()
	default:
		frontendURL, serverURL, cleanup = startFrontAndBackend()
	}
	defer cleanup()

	exitCode := m.Run()

	os.Exit(exitCode)
}

// withTestContext executes the provided test function with different test contexts
// based on environment variables to control which tests run
func withTestContext(t *testing.T, testFn func(t *testing.T, ctx *testContext)) {
	runApplication := os.Getenv("TEST_TYPE") == "application" || os.Getenv("TEST_TYPE") == ""
	runBackEnd := os.Getenv("TEST_TYPE") == "back-end" || os.Getenv("TEST_TYPE") == ""
	runFrontEnd := os.Getenv("TEST_TYPE") == "front-end" || os.Getenv("TEST_TYPE") == ""

	if runApplication {
		t.Run("Application", func(t *testing.T) {
			ctx := newTestContext(testhelpers.NewDomainTestDriver())
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}

	if runBackEnd {
		t.Run("HTTPExecutable", func(t *testing.T) {
			httpDriver := httpdriver.New(serverURL)
			ctx := newTestContext(httpDriver)
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}

	if runFrontEnd {
		t.Run("FrontEnd", func(t *testing.T) {
			uiDriver := uidriver.New(t, frontendURL)

			ctx := newTestContext(uiDriver)
			// Clear at start to ensure clean state
			ctx.clearAll()
			// Extra wait to ensure clear fully propagates before test starts
			time.Sleep(1 * time.Second)
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}
}
