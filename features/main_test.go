package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/features"
	"github.com/sirockin/cucumber-screenplay-go/features/driver/domain"
	httpdriver "github.com/sirockin/cucumber-screenplay-go/features/driver/http"
)

func TestDomainFeatures(t *testing.T){
	features.Test(t, domain.New(), []string{"."})
}

func TestHTTPFeatures(t *testing.T){
	// Note: This test requires an HTTP server implementing the API to be running on localhost:8080
	// In a real scenario, you would either:
	// 1. Start a test server programmatically
	// 2. Use docker-compose or similar for integration tests
	// 3. Skip this test if server is not available
	t.Skip("HTTP API server not running - implement server first or run manually")

	httpClient := httpdriver.New("http://localhost:8080")
	features.Test(t, httpClient, []string{"."})
}
