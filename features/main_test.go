package features_test

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/sirockin/cucumber-screenplay-go/features/driver"
	"github.com/sirockin/cucumber-screenplay-go/features/driver/domain"
	httpdriver "github.com/sirockin/cucumber-screenplay-go/features/driver/http"
	internaldomain "github.com/sirockin/cucumber-screenplay-go/internal/domain"
	httpserver "github.com/sirockin/cucumber-screenplay-go/internal/http"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestDomain(t *testing.T) {
	RunSuite(t, domain.New(), []string{"."})
}

// TestHTTPInProcess tests against an in-process HTTP server
func TestHTTPInProcess(t *testing.T) {
	// Start test server and get its URL
	serverURL := startInProcessServer(t, domain.New())

	// Create HTTP client pointing to our test server
	httpClient := httpdriver.New(serverURL)

	// Run the same BDD tests against the HTTP API
	RunSuite(t, httpClient, []string{"."})
}

// TestHttpExecutable tests against the actual running server executable
func TestHttpExecutable(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Start real server executable and get its URL
	serverURL := startServerExecutable(t)

	// Create HTTP client pointing to the real server
	httpClient := httpdriver.New(serverURL)

	// Run the same BDD tests against the actual server executable
	RunSuite(t, httpClient, []string{"."})
}

// TestHttpDocker tests against the server running in a Docker container using testcontainers
func TestHttpDocker(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Docker integration test in short mode")
	}

	// Check if Docker is available
	if !isDockerAvailable(t) {
		t.Skip("Docker not available, skipping Docker integration test")
	}

	// Build and start Docker container using testcontainers
	serverURL := startTestContainer(t)

	// Create HTTP client pointing to the containerized server
	httpClient := httpdriver.New(serverURL)

	// Run the same BDD tests against the containerized server
	RunSuite(t, httpClient, []string{"."})
}

// startServerExecutable builds and starts the actual server executable
// and returns the server URL. Cleanup is handled automatically via t.Cleanup.
func startServerExecutable(t *testing.T) string {
	// Build the server executable
	serverBinary := buildServerExecutable(t)

	// Find an available port
	port := findAvailablePort(t)

	// Start the server process
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, serverBinary, "-port", strconv.Itoa(port))

	// Set up process group for clean termination
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Capture server output for debugging
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("Failed to create stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		t.Fatalf("Failed to create stderr pipe: %v", err)
	}

	// Start the server
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Monitor server output in background
	go logServerOutput(t, "STDOUT", stdout)
	go logServerOutput(t, "STDERR", stderr)

	// Wait for server to be ready
	serverURL := fmt.Sprintf("http://localhost:%d", port)
	waitForServerReady(t, serverURL)

	t.Logf("Server started successfully at %s (PID: %d)", serverURL, cmd.Process.Pid)

	// Register cleanup function
	t.Cleanup(func() {
		t.Logf("Shutting down server (PID: %d)", cmd.Process.Pid)

		// Cancel context to stop the command
		cancel()

		// Give the process time to shut down gracefully
		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()

		select {
		case <-done:
			t.Logf("Server shut down gracefully")
		case <-time.After(5 * time.Second):
			t.Logf("Server didn't shut down gracefully, killing process group")
			// Kill the entire process group to ensure cleanup
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			<-done
		}

		// Clean up the binary
		os.Remove(serverBinary)
	})

	return serverURL
}

// buildServerExecutable builds the server and returns the path to the executable
func buildServerExecutable(t *testing.T) string {
	// Create temporary executable path
	tempDir := t.TempDir()
	serverBinary := filepath.Join(tempDir, "test-server")

	// Build the server
	t.Logf("Building server executable...")
	cmd := exec.Command("go", "build", "-o", serverBinary, "./cmd/server")

	// Set working directory to project root (parent of features)
	cmd.Dir = ".."

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build server: %v\nOutput: %s", err, output)
	}

	t.Logf("Server built successfully at %s", serverBinary)
	return serverBinary
}

// findAvailablePort finds an available port for the test server
func findAvailablePort(t *testing.T) int {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to find available port: %v", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()
	return port
}

// waitForServerReady waits for the server to be ready to accept connections
func waitForServerReady(t *testing.T, serverURL string) {
	client := &http.Client{Timeout: 1 * time.Second}

	// Try to connect to the server health endpoint
	for i := range 30 { // Wait up to 30 seconds
		resp, err := client.Get(serverURL + "/accounts")
		if err == nil {
			resp.Body.Close()
			// Even a 405 Method Not Allowed means the server is responding
			if resp.StatusCode == http.StatusMethodNotAllowed || resp.StatusCode < 500 {
				t.Logf("Server is ready after %d attempts", i+1)
				return
			}
		}

		time.Sleep(1 * time.Second)
	}

	t.Fatalf("Server did not become ready within 30 seconds")
}

// logServerOutput logs server output for debugging
func logServerOutput(t *testing.T, prefix string, pipe io.ReadCloser) {
	defer pipe.Close()
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		line := scanner.Text()
		// Only log important messages to avoid test output noise
		if strings.Contains(line, "Starting server") ||
			strings.Contains(line, "error") ||
			strings.Contains(line, "Error") ||
			strings.Contains(line, "failed") ||
			strings.Contains(line, "Failed") {
			t.Logf("[SERVER %s] %s", prefix, line)
		}
	}
}

// startInProcessServer starts an HTTP server wrapping the given ApplicationDriver
// and returns the server URL. Cleanup is handled automatically via t.Cleanup.
func startInProcessServer(t *testing.T, app driver.ApplicationDriver) string {
	// Extract the domain from the test driver
	var domainInstance *internaldomain.Domain
	if testDriver, ok := app.(interface{ Domain() *internaldomain.Domain }); ok {
		domainInstance = testDriver.Domain()
	} else {
		t.Fatalf("ApplicationDriver does not provide access to domain instance")
	}

	// Create HTTP server using internal implementation directly
	server := httpserver.NewServer(domainInstance)

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

// isDockerAvailable checks if Docker is available on the system
func isDockerAvailable(t *testing.T) bool {
	_, err := runCommand(t, "docker", "version")
	return err == nil
}

// startTestContainer builds and starts a Docker container using testcontainers
// and returns the server URL. Cleanup is handled automatically via testcontainers.
func startTestContainer(t *testing.T) string {
	ctx := context.Background()

	// Get absolute path to project root
	projectRoot, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("Failed to get project root path: %v", err)
	}

	// Create container request using the deploy/Dockerfile
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    projectRoot,
			Dockerfile: "./deploy/Dockerfile",
		},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForLog("API endpoints"),
	}

	// Start container
	t.Logf("Starting testcontainer...")
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %v", err)
	}

	// Clean up container
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Warning: Failed to terminate container: %v", err)
		}
	})

	// Get mapped port
	mappedPort, err := container.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatalf("Failed to get mapped port: %v", err)
	}

	// Get host
	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}

	serverURL := fmt.Sprintf("http://%s:%s", host, mappedPort.Port())
	t.Logf("Container started successfully at %s", serverURL)

	return serverURL
}

// Helper functions for command execution

func runCommand(t *testing.T, name string, args ...string) ([]byte, error) {
	return runCommandWithTimeout(t, 30*time.Second, name, args...)
}

func runCommandWithTimeout(_ *testing.T, timeout time.Duration, name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.CombinedOutput()
}
