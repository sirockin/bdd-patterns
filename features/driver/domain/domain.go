package domain

import (
	"github.com/sirockin/cucumber-screenplay-go/internal/domain"
)

// New creates a new domain test driver that wraps the actual domain
func New() *TestDriver {
	return &TestDriver{
		domain: domain.New(),
	}
}

// TestDriver is a test driver that delegates to the actual domain
type TestDriver struct {
	domain *domain.Domain
}

// Domain returns the underlying domain for direct access
func (t *TestDriver) Domain() *domain.Domain {
	return t.domain
}

func (t *TestDriver) ClearAll() {
	t.domain.ClearAll()
}

func (t *TestDriver) CreateAccount(name string) error {
	return t.domain.CreateAccount(name)
}

func (t *TestDriver) GetAccount(name string) (domain.Account, error) {
	return t.domain.GetAccount(name)
}

func (t *TestDriver) Activate(name string) error {
	return t.domain.Activate(name)
}

func (t *TestDriver) IsActivated(name string) bool {
	return t.domain.IsActivated(name)
}

func (t *TestDriver) Authenticate(name string) error {
	return t.domain.Authenticate(name)
}

func (t *TestDriver) IsAuthenticated(name string) bool {
	return t.domain.IsAuthenticated(name)
}

func (t *TestDriver) GetProjects(name string) ([]domain.Project, error) {
	return t.domain.GetProjects(name)
}

func (t *TestDriver) CreateProject(name string) error {
	return t.domain.CreateProject(name)
}
