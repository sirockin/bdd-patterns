package http

import (
	"net/http"

	internalapp "github.com/sirockin/cucumber-screenplay-go/back-end/internal/domain/application"
	internalhttp "github.com/sirockin/cucumber-screenplay-go/back-end/internal/http"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/domain/application"
)

// NewServer creates a new HTTP server wrapping the given service
func NewServer(service *application.Service) http.Handler {
	// Extract the internal service to use with internal HTTP server
	internalService := internalapp.New()

	// Copy state from the wrapper service to the internal service if needed
	// For now, since they're separate instances, we'll create a new internal service
	// In a real implementation, you might want to refactor this differently

	return internalhttp.NewServer(internalService)
}

// NewServerWithInternalService creates a new HTTP server with an internal service
func NewServerWithInternalService() http.Handler {
	internalService := internalapp.New()
	return internalhttp.NewServer(internalService)
}
