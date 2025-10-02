package features_test

import (
	"testing"

	httpdriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/http"
	uidriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/ui"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"
)

func TestDomain(t *testing.T) {
	runTests(t, testhelpers.NewDomainTestDriver())
}

// TestBackEnd tests against the actual running server executable
func TestBackEnd(t *testing.T) {
	serverURL := startServerExecutable(t)

	httpDriver := httpdriver.New(serverURL)

	runTests(t, httpDriver)
}

// TestFrontEnd tests against both frontend and API running in containers using UI automation
func TestFrontEnd(t *testing.T) {
	frontendURL := startFrontAndBackend(t)

	uiDriver := uidriver.New(t, frontendURL)

	runTests(t, uiDriver)
}
