package testhelpers

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/sirockin/cucumber-screenplay-go/back-end/internal/domain/application"
	httpserver "github.com/sirockin/cucumber-screenplay-go/back-end/internal/http"
)

// Create an in-process server for testing
func NewInProcessServer(t *testing.T) string {
	// Create HTTP server using internal implementation directly
	server := httpserver.NewServer(application.New())

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

	// Register cleanup function and return server URL
	serverURL := fmt.Sprintf("http://localhost:%d", port)
	t.Cleanup(func() {
		httpServer.Close()
	})

	return serverURL
}
