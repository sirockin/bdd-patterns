package driver

import (
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/entities"
)

// TestDriver is our interface to the system under test
type TestDriver interface {
	CreateAccount(name string) error
	ClearAll()
	GetAccount(name string) (entities.Account, error)
	Authenticate(name string) error
	IsAuthenticated(name string) bool
	Activate(name string) error
	CreateProject(name string) error
	GetProjects(name string) ([]entities.Project, error)
}
