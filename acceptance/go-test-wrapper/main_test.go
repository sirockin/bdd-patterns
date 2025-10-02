package features_test

import (
	"os"
	"testing"

	httpdriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/http"
	uidriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/ui"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"
)

var (
	serverURL   string
	frontendURL string
)

func TestMain(m *testing.M) {
	var cleanupFuncs []func()

	if os.Getenv("RUN_BACKEND") == "true" {
		var cleanup func()
		serverURL, cleanup = startServerExecutable()
		cleanupFuncs = append(cleanupFuncs, cleanup)
	} else if os.Getenv("RUN_FRONTEND") == "true" {
		var cleanup func()
		frontendURL, cleanup = startFrontAndBackend()
		cleanupFuncs = append(cleanupFuncs, cleanup)
	}

	exitCode := m.Run()

	// Run all cleanup functions
	for _, cleanup := range cleanupFuncs {
		cleanup()
	}

	os.Exit(exitCode)
}

// withTestContext executes the provided test function with different test contexts
// based on environment variables to control which tests run
func withTestContext(t *testing.T, testFn func(t *testing.T, ctx *testContext)) {
	if os.Getenv("RUN_APPLICATION") == "true" {
		t.Run("Application", func(t *testing.T) {
			ctx := newTestContext(testhelpers.NewDomainTestDriver())
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}

	if os.Getenv("RUN_BACK_END") == "true" {
		t.Run("HTTPExecutable", func(t *testing.T) {
			httpDriver := httpdriver.New(serverURL)
			ctx := newTestContext(httpDriver)
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}

	if os.Getenv("RUN_FRONT_END") == "true" {
		t.Run("FrontEnd", func(t *testing.T) {
			uiDriver := uidriver.New(t, frontendURL)

			ctx := newTestContext(uiDriver)
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}
}
