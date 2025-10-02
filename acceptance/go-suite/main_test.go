package features_test

import (
	"testing"

	httpdriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/http"
	uidriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/ui"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"
	"github.com/stretchr/testify/suite"
)

func TestDomain(t *testing.T) {
	suite.Run(t, NewFeatureSuite(testhelpers.NewDomainTestDriver()))
}

// TestBackEnd tests against the actual running server executable
func TestBackEnd(t *testing.T) {
	serverURL := startServerExecutable(t)

	httpDriver := httpdriver.New(serverURL)

	suite.Run(t, NewFeatureSuite(httpDriver))
}

// TestFrontEnd tests against both frontend and API running in containers using UI automation
func TestFrontEnd(t *testing.T) {
	frontendURL := startFrontAndBackend(t)

	uiDriver := uidriver.New(t, frontendURL)

	suite.Run(t, NewFeatureSuite(uiDriver))
}

func TestBackAndFrontEndDocker(t *testing.T) {
	backendURL, frontendURL := startFrontAndBackendDocker(t)

	t.Run("BackEnd", func(t *testing.T) {
		httpDriver := httpdriver.New(backendURL)
		suite.Run(t, NewFeatureSuite(httpDriver))
	})

	t.Run("FrontEnd", func(t *testing.T) {
		uiDriver := uidriver.New(t, frontendURL)
		suite.Run(t, NewFeatureSuite(uiDriver))
	})
}
