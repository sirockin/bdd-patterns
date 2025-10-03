package features_test

import (
	"os"
	"testing"
	"time"
)

var serverURL string

func TestMain(m *testing.M) {
	var cleanup func()
	serverURL, cleanup = startServerExecutable()
	defer cleanup()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func setupTest(t *testing.T) *testContext {
	ctx := newTestContext(t, serverURL)
	// Clear at start to ensure clean state
	ctx.clearAll()
	// Extra wait to ensure clear fully propagates before test starts
	time.Sleep(1 * time.Second)
	t.Cleanup(func() {
		ctx.clearAll()
	})
	return ctx
}
