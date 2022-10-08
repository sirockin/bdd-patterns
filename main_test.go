package main_test

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/sirockin/cucumber-screenplay-go/domain"
)

// Application is our interface to the system under test
type Application interface{
	CreateAccount(name string)error
	ClearAll()
	GetAccount(name string)(domain.Account,error)
	Authenticate(name string)error
	IsAuthenticated(name string)bool
	Activate(name string)error
	CreateProject(name string)error
	GetProjects(name string)([]domain.Project,error)
}


///////////////////////
// Actions

var CreateAccount = struct {
	forThemselves Action
}{
	forThemselves: func(abilities Abilities)error{
		return abilities.app.CreateAccount(abilities.name)
	},
}

var Activate = struct {
	theirAccount Action
}{
	theirAccount: func(abilities Abilities)error{
		if _, err := abilities.app.GetAccount(abilities.name); err != nil{
			return nil
		} 
		return abilities.app.Activate(abilities.name)
	},
}

func signUp(abilities Abilities)error{
	return abilities.attemptsTo(
		CreateAccount.forThemselves,
		Activate.theirAccount,
	)
} 

func signIn(abilities Abilities)error{
	return abilities.app.Authenticate(abilities.name)
}

func createProject(abilities Abilities)error{
	return abilities.app.CreateProject(abilities.name)
}

// Questions
func amIAuthenticated(abilities Abilities)(interface{},error){
	return abilities.app.IsAuthenticated(abilities.name), nil
}

func howManyProjectsDoIHave(abilities Abilities)(interface{},error){
	projects, err := abilities.app.GetProjects(abilities.name)
	if err != nil {
		return 0, err
	}
	return len(projects), nil
}

type accountFeature struct {
	actors map[string]*Actor
	app Application
}

func(a *accountFeature) reset(){
	a.actors = make(map[string]*Actor)
	a.app.ClearAll()
}

func(a *accountFeature) Actor(name string)*Actor{
	if a.actors[name]==nil {
		a.actors[name]=NewActor(name, a.app)
	}
	return a.actors[name]	
}

////////
// Steps

func (a *accountFeature) personHasCreatedAnAccount(name string) error {
	return a.Actor(name).attemptsTo(CreateAccount.forThemselves)
}

func (a *accountFeature) personHasSignedUp(name string) error {
	return a.Actor(name).attemptsTo(signUp)
}

func (a *accountFeature) personShouldBeAuthenticated(name string) error {
	return a.Actor(name).expectsAnswer(amIAuthenticated,true)
}

func (a *accountFeature) personShouldNotBeAuthenticated(name string) error {
	return a.Actor(name).expectsAnswer(amIAuthenticated,false)
}

func (a *accountFeature) personShouldNotSeeAnyProjects(name string) error {
	return a.Actor(name).expectsAnswer(howManyProjectsDoIHave,0)
}

func (a *accountFeature) personShouldSeeTheirProject(name string) error {
	return a.Actor(name).expectsAnswer(howManyProjectsDoIHave,1)
}

func (a *accountFeature) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	return a.Actor(name).expectsLastErrorToContain("you need to activate your account")
}

func (a *accountFeature) personTriesToSignIn(name string) error {
	a.Actor(name).attemptsTo(signIn)
	return nil;	// The step succeeds even if the result is bad to allow the next step to check the error
}

func (a *accountFeature) personCreatesAProject(name string) error {
	return a.Actor(name).attemptsTo(createProject)
}

func (a *accountFeature) personActivatesTheirAccount(name string) error {
	return a.Actor(name).attemptsTo(Activate.theirAccount)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
	  ScenarioInitializer: InitializeScenario,
	  Options: &godog.Options{
		Format:   "pretty",
		Paths:    []string{"features"},
		TestingT: t, // Testing instance that will run subtests.
	  },
	}
  
	if suite.Run() != 0 {
	  t.Fatal("non-zero status returned, failed to run feature tests")
	}
  }

func InitializeScenario(ctx *godog.ScenarioContext) {
	af := &accountFeature{
		app: domain.New(),
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {		
		af.reset()
		return ctx, nil
	})

	ctx.Step(`^(Bob|Tanya|Sue) has created an account$`, af.personHasCreatedAnAccount)
	ctx.Step(`^(Bob|Tanya|Sue) has signed up$`, af.personHasSignedUp)
	ctx.Step(`^(Bob|Tanya|Sue) should not be authenticated$`, af.personShouldNotBeAuthenticated)
	ctx.Step(`^(Bob|Tanya|Sue) should not see any projects$`, af.personShouldNotSeeAnyProjects)
	ctx.Step(`^(Bob|Tanya|Sue) should see an error telling (him|her|them) to activate the account$`, af.personShouldSeeAnErrorTellingThemToActivateTheAccount)
	ctx.Step(`^(Bob|Tanya|Sue) tries to sign in$`, af.personTriesToSignIn)
	ctx.Step(`^(Bob|Tanya|Sue) creates a project$`, af.personCreatesAProject)
	ctx.Step(`^(Bob|Tanya|Sue) should see (his|her|the) project$`, af.personShouldSeeTheirProject)
	ctx.Step(`^(Bob|Tanya|Sue) activates (his|her) account$`, af.personActivatesTheirAccount)
	ctx.Step(`^(Bob|Tanya|Sue) should be authenticated$`, af.personShouldBeAuthenticated)
}
