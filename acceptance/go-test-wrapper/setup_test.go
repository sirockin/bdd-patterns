package features_test

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// startServerExecutable builds and starts the actual server executable using the root makefile
// and returns the server URL and a cleanup function.
func startServerExecutable() (string, func()) {
	// Start the server process using root makefile target
	cmd := exec.Command("make", "run-backend")

	// Set working directory to project root
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		log.Printf("Failed to get project root path: %v", err)
		os.Exit(1)
	}
	cmd.Dir = projectRoot

	// Set up process group for clean termination
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Capture server output for debugging
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Failed to create stdout pipe: %v", err)
		os.Exit(1)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("Failed to create stderr pipe: %v", err)
		os.Exit(1)
	}

	// Start the server
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start server: %v", err)
		os.Exit(1)
	}

	// Monitor server output in background
	go logServerOutput("STDOUT", stdout)
	go logServerOutput("STDERR", stderr)

	// Wait for server to be ready
	serverURL := "http://localhost:8080"
	waitForServerReady(serverURL)

	log.Printf("Server started successfully at %s (PID: %d)", serverURL, cmd.Process.Pid)

	// Create cleanup function
	cleanup := func() {
		log.Printf("Shutting down server (PID: %d)", cmd.Process.Pid)

		// Kill the entire process group to ensure cleanup
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)

		// Give the process time to shut down
		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()

		select {
		case <-done:
			log.Printf("Server shut down gracefully")
		case <-time.After(5 * time.Second):
			log.Printf("Server didn't shut down gracefully")
		}
	}

	return serverURL, cleanup
}

// waitForServerReady waits for the server to be ready to accept connections
func waitForServerReady(serverURL string) {
	waitForServerReadyWithTimeout(serverURL, 30*time.Second)
}

// waitForServerReadyWithTimeout waits for the server to be ready with a custom timeout
func waitForServerReadyWithTimeout(serverURL string, timeout time.Duration) {
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
				log.Printf("Server is ready at %s after %d attempts (%.1fs)", serverURL, attempt, time.Since(deadline.Add(-timeout)).Seconds())
				return
			}
		}

		time.Sleep(2 * time.Second)
	}

	log.Printf("Server at %s did not become ready within %v (tried %d times)", serverURL, timeout, attempt)
	os.Exit(1)
}

// logServerOutput logs server output for debugging
func logServerOutput(prefix string, pipe io.ReadCloser) {
	defer pipe.Close()
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		line := scanner.Text()
		// Log all output to help diagnose build and startup issues
		log.Printf("[SERVER %s] %s", prefix, line)
	}
}

// Start back end and front end services by calling `make run` and return the front end URL and cleanup function
func startFrontAndBackend() (string, func()) {
	// Start both back end and front end services
	cmd := exec.Command("make", "run")

	// Set working directory to project root
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		log.Printf("Failed to get project root path: %v", err)
		os.Exit(1)
	}
	cmd.Dir = projectRoot

	// Set up process group for clean termination
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Capture output for debugging
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Failed to create stdout pipe: %v", err)
		os.Exit(1)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("Failed to create stderr pipe: %v", err)
		os.Exit(1)
	}

	// Start the services
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start services: %v", err)
		os.Exit(1)
	}

	// Monitor output in background
	go logServerOutput("FRONTEND STDOUT", stdout)
	go logServerOutput("FRONTEND STDERR", stderr)

	// Wait for frontend to be ready (dev server takes longer to compile and start)
	frontendURL := "http://localhost:3000"
	waitForServerReadyWithTimeout(frontendURL, 60*time.Second)

	log.Printf("Frontend started successfully at %s (PID: %d)", frontendURL, cmd.Process.Pid)

	// Create cleanup function
	cleanup := func() {
		log.Printf("Shutting down services (PID: %d)", cmd.Process.Pid)

		// Kill the entire process group to ensure cleanup
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)

		// Give the process time to shut down gracefully
		done := make(chan error, 1)
		go func() {
			done <- cmd.Wait()
		}()

		select {
		case <-done:
			log.Printf("Services shut down gracefully")
		case <-time.After(5 * time.Second):
			log.Printf("Services didn't shut down gracefully, killing process group")
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			<-done
		}
	}
	return frontendURL, cleanup
}
