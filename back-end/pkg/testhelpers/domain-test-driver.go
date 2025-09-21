package testhelpers

import (
	"github.com/sirockin/cucumber-screenplay-go/back-end/internal/domain/application"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/entities"
)

// New creates a new acceptance test driver that wraps the actual domain
func NewDomainTestDriver() *DomainTestDriver {
	return &DomainTestDriver{
		appService: application.New(),
	}
}

// DomainTestDriver is a test driver that delegates to the actual domain
// It implements the AcceptanceTestDriver interface implicitly
type DomainTestDriver struct {
	appService *application.Service
}

func (t *DomainTestDriver) ClearAll() {
	t.appService.ClearAll()
}

func (t *DomainTestDriver) CreateAccount(name string) error {
	return t.appService.CreateAccount(name)
}

func (t *DomainTestDriver) GetAccount(name string) (entities.Account, error) {
	return t.appService.GetAccount(name)
}

func (t *DomainTestDriver) Activate(name string) error {
	return t.appService.Activate(name)
}

func (t *DomainTestDriver) IsActivated(name string) bool {
	return t.appService.IsActivated(name)
}

func (t *DomainTestDriver) Authenticate(name string) error {
	return t.appService.Authenticate(name)
}

func (t *DomainTestDriver) IsAuthenticated(name string) bool {
	return t.appService.IsAuthenticated(name)
}

func (t *DomainTestDriver) GetProjects(name string) ([]entities.Project, error) {
	return t.appService.GetProjects(name)
}

func (t *DomainTestDriver) CreateProject(name string) error {
	return t.appService.CreateProject(name)
}