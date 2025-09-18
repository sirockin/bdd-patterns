package http

import (
	"github.com/sirockin/cucumber-screenplay-go/features/driver"
	httpserver "github.com/sirockin/cucumber-screenplay-go/internal/http"
)

// NewServer creates a new HTTP server wrapping the given ApplicationDriver
func NewServer(app driver.ApplicationDriver) *httpserver.Server {
	return httpserver.NewServer(app)
}