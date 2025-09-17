package http

import (
	"github.com/sirockin/cucumber-screenplay-go/features/driver"
	"github.com/sirockin/cucumber-screenplay-go/internal/server"
)

// NewServer creates a new HTTP server wrapping the given ApplicationDriver
func NewServer(app driver.ApplicationDriver) *server.Server {
	return server.NewServer(app)
}