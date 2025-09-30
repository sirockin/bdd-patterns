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

	httpdriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/http"
	uidriver "github.com/sirockin/cucumber-screenplay-go/acceptance/driver/ui"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

// withTestContext executes the provided test function with different test contexts
// based on environment variables to control which tests run
func withTestContext(t *testing.T, testFn func(t *testing.T, ctx *testContext)) {
	// Check which drivers to run based on environment variables
	runApplication := os.Getenv("RUN_APPLICATION") != "false"
	runHTTP := os.Getenv("RUN_HTTP") != "false"
	runUI := os.Getenv("RUN_UI") != "false"

	// Default behavior: if no specific environment is set, run all non-slow tests
	if os.Getenv("RUN_APPLICATION") == "" &&
		os.Getenv("RUN_HTTP") == "" &&
		os.Getenv("RUN_UI") == "" {
		runApplication = true
		runHTTP = !testing.Short()
		runUI = !testing.Short() && isDockerAvailable(t)
	}

	if runApplication {
		t.Run("Application", func(t *testing.T) {
			ctx := newTestContext(testhelpers.NewDomainTestDriver())
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}

	if runHTTP {
		t.Run("HTTP", func(t *testing.T) {
			if testing.Short() {
				t.Skip("Skipping integration test in short mode")
			}

			serverURL := startServerExecutable(t)
			httpDriver := httpdriver.New(serverURL)
			ctx := newTestContext(httpDriver)
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}

	if runUI {
		t.Run("UI", func(t *testing.T) {
			if testing.Short() {
				t.Skip("Skipping UI integration test in short mode")
			}

			if !isDockerAvailable(t) {
				t.Skip("Docker not available, skipping UI integration test")
			}

			frontendURL := startUITestEnvironment(t)
			uiDriver, err := uidriver.New(frontendURL)
			if err != nil {
				t.Fatalf("Failed to create UI driver: %v", err)
			}

			t.Cleanup(func() {
				if err := uiDriver.Close(); err != nil {
					t.Logf("Warning: Failed to close UI driver: %v", err)
				}
			})

			ctx := newTestContext(uiDriver)
			t.Cleanup(func() {
				ctx.clearAll()
			})
			testFn(t, ctx)
		})
	}
}

// startServerExecutable builds and starts the actual server executable
// and returns the server URL. Cleanup is handled automatically via t.Cleanup.
func startServerExecutable(t *testing.T) string {
	serverBinary := buildServerExecutable(t)
	port := findAvailablePort(t)

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, serverBinary, "-port", strconv.Itoa(port))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("Failed to create stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		t.Fatalf("Failed to create stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	go logServerOutput(t, "STDOUT", stdout)
	go logServerOutput(t, "STDERR", stderr)

	serverURL := fmt.Sprintf("http://localhost:%d", port)
	waitForServerReady(t, serverURL)

	t.Logf("Server started successfully at %s (PID: %d)", serverURL, cmd.Process.Pid)

	t.Cleanup(func() {
		t.Logf("Shutting down server (PID: %d)", cmd.Process.Pid)
		cancel()

		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()

		select {
		case <-done:
			t.Logf("Server shut down gracefully")
		case <-time.After(5 * time.Second):
			t.Logf("Server didn't shut down gracefully, killing process group")
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			<-done
		}

		os.Remove(serverBinary)
	})

	return serverURL
}

func buildServerExecutable(t *testing.T) string {
	tempDir := t.TempDir()
	serverBinary := filepath.Join(tempDir, "test-server")

	t.Logf("Building server executable...")
	cmd := exec.Command("go", "build", "-o", serverBinary, "./cmd/server")
	cmd.Dir = "../../back-end"

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build server: %v\nOutput: %s", err, output)
	}

	t.Logf("Server built successfully at %s", serverBinary)
	return serverBinary
}

func findAvailablePort(t *testing.T) int {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to find available port: %v", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()
	return port
}

func waitForServerReady(t *testing.T, serverURL string) {
	client := &http.Client{Timeout: 1 * time.Second}

	for i := range 30 {
		resp, err := client.Get(serverURL + "/accounts")
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusMethodNotAllowed || resp.StatusCode < 500 {
				t.Logf("Server is ready after %d attempts", i+1)
				return
			}
		}
		time.Sleep(1 * time.Second)
	}

	t.Fatalf("Server did not become ready within 30 seconds")
}

func logServerOutput(t *testing.T, prefix string, pipe io.ReadCloser) {
	defer pipe.Close()
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Starting server") ||
			strings.Contains(line, "error") ||
			strings.Contains(line, "Error") ||
			strings.Contains(line, "failed") ||
			strings.Contains(line, "Failed") {
			t.Logf("[SERVER %s] %s", prefix, line)
		}
	}
}

func isDockerAvailable(t *testing.T) bool {
	_, err := runCommand(t, "docker", "version")
	return err == nil
}

func runCommand(t *testing.T, name string, args ...string) ([]byte, error) {
	return runCommandWithTimeout(t, 30*time.Second, name, args...)
}

func runCommandWithTimeout(_ *testing.T, timeout time.Duration, name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.CombinedOutput()
}

func startUITestEnvironment(t *testing.T) string {
	ctx := context.Background()

	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("Failed to get project root path: %v", err)
	}

	nw, err := network.New(ctx)
	if err != nil {
		t.Fatalf("Failed to create network: %v", err)
	}

	t.Cleanup(func() {
		if err := nw.Remove(ctx); err != nil {
			t.Logf("Warning: Failed to remove network: %v", err)
		}
	})

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

	t.Cleanup(func() {
		if err := apiContainer.Terminate(ctx); err != nil {
			t.Logf("Warning: Failed to terminate API container: %v", err)
		}
	})

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

	t.Cleanup(func() {
		if err := frontendContainer.Terminate(ctx); err != nil {
			t.Logf("Warning: Failed to terminate frontend container: %v", err)
		}
	})

	frontendPort, err := frontendContainer.MappedPort(ctx, "80")
	if err != nil {
		t.Fatalf("Failed to get frontend mapped port: %v", err)
	}

	frontendHost, err := frontendContainer.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get frontend container host: %v", err)
	}

	frontendURL := fmt.Sprintf("http://%s:%s", frontendHost, frontendPort.Port())
	t.Logf("UI test environment started successfully, frontend at %s", frontendURL)

	return frontendURL
}
