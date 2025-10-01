package features_test

import (
	"testing"

	httpdriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/http"
	uidriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/ui"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"
)

func TestDomain(t *testing.T) {
	RunSuite(t, testhelpers.NewDomainTestDriver(), []string{"."})
}

// TestBackEnd tests against the actual running server executable
func TestBackEnd(t *testing.T) {
	// Start real server executable and get its URL
	serverURL := startServerExecutable(t)

	// Create HTTP driver pointing to the real server
	httpDriver := httpdriver.New(serverURL)

	RunSuite(t, httpDriver, []string{"."})
}

// TestFrontEnd tests against both frontend and API running in containers using UI automation
func TestFrontEnd(t *testing.T) {
	// Start both frontend and API containers with UI test driver
	frontendURL := startFrontAndBackend(t)

	// Create UI driver pointing to the frontend
	uiDriver, err := uidriver.New(frontendURL)
	if err != nil {
		t.Fatalf("Failed to create UI driver: %v", err)
	}

	// Ensure cleanup of browser resources
	t.Cleanup(func() {
		if err := uiDriver.Close(); err != nil {
			t.Logf("Warning: Failed to close UI driver: %v", err)
		}
	})

	// Run the same BDD tests against the UI
	RunSuite(t, uiDriver, []string{"."})
}

func TestBackAndFrontEndDocker(t *testing.T) {
	// Start both frontend and API containers with UI test driver
	frontendURL := startFrontAndBackendDocker(t)

	// Create UI driver pointing to the frontend
	uiDriver, err := uidriver.New(frontendURL)
	if err != nil {
		t.Fatalf("Failed to create UI driver: %v", err)
	}

	// Ensure cleanup of browser resources
	t.Cleanup(func() {
		if err := uiDriver.Close(); err != nil {
			t.Logf("Warning: Failed to close UI driver: %v", err)
		}
	})

	// Run the same BDD tests against the UI
	RunSuite(t, uiDriver, []string{"."})
}
