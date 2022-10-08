package features

import (
	"github.com/sirockin/cucumber-screenplay-go/domain"
)

// ApplicationDriver is our interface to the system under test
type ApplicationDriver interface{
	CreateAccount(name string)error
	ClearAll()
	GetAccount(name string)(domain.Account,error)
	Authenticate(name string)error
	IsAuthenticated(name string)bool
	Activate(name string)error
	CreateProject(name string)error
	GetProjects(name string)([]domain.Project,error)
}
