package features_test

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/sirockin/cucumber-screenplay-go/features"
	"github.com/sirockin/cucumber-screenplay-go/features/driver"
	"github.com/sirockin/cucumber-screenplay-go/features/driver/domain"
	httpdriver "github.com/sirockin/cucumber-screenplay-go/features/driver/http"
)

func TestDomainFeatures(t *testing.T){
	features.Test(t, domain.New(), []string{"."})
}

func TestHTTPFeatures(t *testing.T){
	// Start test server and get its URL
	serverURL, cleanup := startTestServer(t, domain.New())
	defer cleanup()

	// Create HTTP client pointing to our test server
	httpClient := httpdriver.New(serverURL)

	// Run the same BDD tests against the HTTP API
	features.Test(t, httpClient, []string{"."})
}

// startTestServer starts an HTTP server wrapping the given ApplicationDriver
// and returns the server URL and a cleanup function
func startTestServer(t *testing.T, app driver.ApplicationDriver) (string, func()) {
	// Create HTTP server wrapping the application
	server := httpdriver.NewServer(app)

	// Find an available port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to find available port: %v", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()

	// Start the HTTP server in a goroutine
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: server,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("HTTP server failed: %v", err)
		}
	}()

	// Wait a moment for server to start
	time.Sleep(100 * time.Millisecond)

	// Return server URL and cleanup function
	serverURL := fmt.Sprintf("http://localhost:%d", port)
	cleanup := func() {
		httpServer.Close()
	}

	return serverURL, cleanup
}
