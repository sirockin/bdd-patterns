package features_test

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

// startServerExecutable builds and starts the actual server executable using the root makefile
// and returns the server URL. Cleanup is handled automatically via t.Cleanup.
func startServerExecutable(t *testing.T) string {
	// Start the server process using root makefile target
	cmd := exec.Command("make", "run-backend")

	// Set working directory to project root
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("Failed to get project root path: %v", err)
	}
	cmd.Dir = projectRoot

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
	serverURL := "http://localhost:8080"
	waitForServerReady(t, serverURL)

	t.Logf("Server started successfully at %s (PID: %d)", serverURL, cmd.Process.Pid)

	// Register cleanup function
	t.Cleanup(func() {
		t.Logf("Shutting down server (PID: %d)", cmd.Process.Pid)

		// Kill the entire process group to ensure cleanup
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)

		// Give the process time to shut down
		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()

		select {
		case <-done:
			t.Logf("Server shut down gracefully")
		case <-time.After(5 * time.Second):
			t.Logf("Server didn't shut down gracefully")
		}
	})

	return serverURL
}

// waitForServerReady waits for the server to be ready to accept connections
func waitForServerReady(t *testing.T, serverURL string) {
	waitForServerReadyWithTimeout(t, serverURL, 30*time.Second)
}

// waitForServerReadyWithTimeout waits for the server to be ready with a custom timeout
func waitForServerReadyWithTimeout(t *testing.T, serverURL string, timeout time.Duration) {
	client := &http.Client{Timeout: 5 * time.Second}

	// Determine which endpoint to check based on URL
	checkPath := "/"
	if strings.Contains(serverURL, ":8080") {
		// Backend server - check /accounts endpoint
		checkPath = "/accounts"
	}

	checkURL := serverURL + checkPath
	deadline := time.Now().Add(timeout)
	attempt := 0

	for time.Now().Before(deadline) {
		attempt++
		resp, err := client.Get(checkURL)
		if err == nil {
			resp.Body.Close()
			// Any response code < 500 means server is responding
			if resp.StatusCode < 500 {
				t.Logf("Server is ready at %s after %d attempts (%.1fs)", serverURL, attempt, time.Since(deadline.Add(-timeout)).Seconds())
				return
			}
		}

		time.Sleep(2 * time.Second)
	}

	t.Fatalf("Server at %s did not become ready within %v (tried %d times)", serverURL, timeout, attempt)
}

// logServerOutput logs server output for debugging
func logServerOutput(t *testing.T, prefix string, pipe io.ReadCloser) {
	defer pipe.Close()
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		line := scanner.Text()
		// Log all output to help diagnose build and startup issues
		t.Logf("[SERVER %s] %s", prefix, line)
	}
}

// isDockerAvailable checks if Docker is available on the system
func isDockerAvailable(t *testing.T) bool {
	_, err := runCommand(t, "docker", "version")
	return err == nil
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

// startUITestEnvironment starts both API and frontend containers using testcontainers
// and returns both backend and frontend URLs. Cleanup is handled automatically via testcontainers.
func startFrontAndBackendDocker(t *testing.T) (backendURL, frontendURL string) {
	if !isDockerAvailable(t) {
		t.Skip("Docker is not available, skipping tests")
	}

	ctx := context.Background()

	// Get absolute path to project root
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("Failed to get project root path: %v", err)
	}

	// Create a network for the containers to communicate
	nw, err := network.New(ctx)
	if err != nil {
		t.Fatalf("Failed to create network: %v", err)
	}

	// Clean up network
	t.Cleanup(func() {
		if err := nw.Remove(ctx); err != nil {
			t.Logf("Warning: Failed to remove network: %v", err)
		}
	})

	// Start API container
	t.Logf("Starting API container...")
	apiContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:    filepath.Join(projectRoot, "back-end"),
				Dockerfile: "Dockerfile",
			},
			ExposedPorts: []string{"8080/tcp"},
			WaitingFor:   wait.ForLog("API endpoints"),
			Networks:     []string{nw.Name},
			NetworkAliases: map[string][]string{
				nw.Name: {"api"},
			},
		},
		Started: true,
	})
	if err != nil {
		t.Fatalf("Failed to start API container: %v", err)
	}

	// Clean up API container
	t.Cleanup(func() {
		if err := apiContainer.Terminate(ctx); err != nil {
			t.Logf("Warning: Failed to terminate API container: %v", err)
		}
	})

	// Get API mapped port
	apiPort, err := apiContainer.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatalf("Failed to get API mapped port: %v", err)
	}

	// Get API host
	apiHost, err := apiContainer.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get API container host: %v", err)
	}

	backendURL = fmt.Sprintf("http://%s:%s", apiHost, apiPort.Port())

	// Start frontend container
	t.Logf("Starting frontend container...")
	frontendContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:    filepath.Join(projectRoot, "front-end"),
				Dockerfile: "Dockerfile",
			},
			ExposedPorts: []string{"80/tcp"},
			WaitingFor:   wait.ForHTTP("/").WithPort("80"),
			Networks:     []string{nw.Name},
			NetworkAliases: map[string][]string{
				nw.Name: {"frontend"},
			},
		},
		Started: true,
	})
	if err != nil {
		t.Fatalf("Failed to start frontend container: %v", err)
	}

	// Clean up frontend container
	t.Cleanup(func() {
		if err := frontendContainer.Terminate(ctx); err != nil {
			t.Logf("Warning: Failed to terminate frontend container: %v", err)
		}
	})

	// Get frontend mapped port
	frontendPort, err := frontendContainer.MappedPort(ctx, "80")
	if err != nil {
		t.Fatalf("Failed to get frontend mapped port: %v", err)
	}

	// Get frontend host
	frontendHost, err := frontendContainer.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get frontend container host: %v", err)
	}

	frontendURL = fmt.Sprintf("http://%s:%s", frontendHost, frontendPort.Port())
	t.Logf("Docker test environment started successfully, backend at %s, frontend at %s", backendURL, frontendURL)

	return backendURL, frontendURL
}

// Start back end and front end services by calling `make run` and return the front end URL
func startFrontAndBackend(t *testing.T) string {
	// Start both back end and front end services
	cmd := exec.Command("make", "run")

	// Set working directory to project root
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("Failed to get project root path: %v", err)
	}
	cmd.Dir = projectRoot

	// Set up process group for clean termination
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Capture output for debugging
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("Failed to create stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		t.Fatalf("Failed to create stderr pipe: %v", err)
	}

	// Start the services
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start services: %v", err)
	}

	// Monitor output in background
	go logServerOutput(t, "FRONTEND STDOUT", stdout)
	go logServerOutput(t, "FRONTEND STDERR", stderr)

	// Wait for frontend to be ready (dev server takes longer to compile and start)
	frontendURL := "http://localhost:3000"
	waitForServerReadyWithTimeout(t, frontendURL, 60*time.Second)

	t.Logf("Frontend started successfully at %s (PID: %d)", frontendURL, cmd.Process.Pid)

	// Register cleanup function
	t.Cleanup(func() {
		t.Logf("Shutting down services (PID: %d)", cmd.Process.Pid)

		// Kill the entire process group to ensure cleanup
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)

		// Give the process time to shut down gracefully
		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()

		select {
		case <-done:
			t.Logf("Services shut down gracefully")
		case <-time.After(5 * time.Second):
			t.Logf("Services didn't shut down gracefully, killing process group")
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			<-done
		}
	})
	return frontendURL
}
