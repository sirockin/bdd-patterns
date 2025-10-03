package features_test

import (
	"os"
	"testing"
	"time"
)

var frontendURL string

func TestMain(m *testing.M) {
	var cleanup func()
	frontendURL, _, cleanup = startFrontAndBackend()
	defer cleanup()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func setupTest(t *testing.T) *testContext {
	ctx := newTestContext(t, frontendURL)
	// Clear at start to ensure clean state
	ctx.clearAll()
	// Extra wait to ensure clear fully propagates before test starts
	time.Sleep(1 * time.Second)
	t.Cleanup(func() {
		ctx.clearAll()
	})
	return ctx
}
