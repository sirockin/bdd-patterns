package application

import (
	"github.com/sirockin/cucumber-screenplay-go/back-end/internal/domain/application"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/domain/entities"
)

// Service wraps the internal application service for external use
type Service struct {
	internal *application.Service
}

// New creates a new service
func New() *Service {
	return &Service{
		internal: application.New(),
	}
}

// ClearAll removes all data
func (s *Service) ClearAll() {
	s.internal.ClearAll()
}

// CreateAccount creates a new account
func (s *Service) CreateAccount(name string) error {
	return s.internal.CreateAccount(name)
}

// GetAccount retrieves an account by name
func (s *Service) GetAccount(name string) (entities.Account, error) {
	return s.internal.GetAccount(name)
}

// Activate activates an account and also authenticates the user
func (s *Service) Activate(name string) error {
	return s.internal.Activate(name)
}

// IsActivated checks if an account is activated
func (s *Service) IsActivated(name string) bool {
	return s.internal.IsActivated(name)
}

// Authenticate authenticates an account (requires activation first)
func (s *Service) Authenticate(name string) error {
	return s.internal.Authenticate(name)
}

// IsAuthenticated checks if an account is authenticated
func (s *Service) IsAuthenticated(name string) bool {
	return s.internal.IsAuthenticated(name)
}

// GetProjects retrieves projects for an account
func (s *Service) GetProjects(name string) ([]entities.Project, error) {
	return s.internal.GetProjects(name)
}

// CreateProject creates a project for an account
func (s *Service) CreateProject(name string) error {
	return s.internal.CreateProject(name)
}