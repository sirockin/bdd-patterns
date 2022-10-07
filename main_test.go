package main

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/sirockin/cucumber-screenplay-go/domain"
)

type Driver interface{
	CreateAccount(name string)error
	ClearAccounts()
	GetAccount(name string)(domain.Account,error)
	Authenticate(name string)error
	IsAuthenticated(name string)bool
	CreateProject(name string)error
	GetProjects(name string)([]domain.Project,error)
}


// Screenplay objects
type Abilities struct {
	name string
	app Driver
	attemptsTo func(actions ...Action)error
}

type Actor struct {
	abilities Abilities
	attemptsTo func(actions ...Action)error
}

type Action func(Abilities)error


func NewActor(name string, app Driver)*Actor{
	ret := &Actor{
		abilities:Abilities{
			name: name, 
			app: app,	
		},
	}
	ret.abilities.attemptsTo=func(actions ...Action)error{
		for i:=0; i<len(actions); i++{
			err := actions[i](ret.abilities)
			if err != nil {
				return err
			}
		}
		return nil
	}
	ret.attemptsTo=ret.abilities.attemptsTo
	return ret
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
		return abilities.app.Authenticate(abilities.name)
	},
}

func signUp(abilities Abilities)error{
	return abilities.attemptsTo(
		CreateAccount.forThemselves,
		Activate.theirAccount,
	)
} 


type accountFeature struct {
	actors map[string]*Actor
	app Driver
}

func(a *accountFeature) reset(){
	a.actors = make(map[string]*Actor)
	a.app.ClearAccounts()
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

func (a *accountFeature) personShouldNotBeAuthenticated(name string) error {
	if a.app.IsAuthenticated(name){
		return fmt.Errorf("Expected %s not to be authenticated", name)
	}
	return nil
}

func (a *accountFeature) personShouldNotSeeAnyProjects(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personTriesToSignIn(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personCreatesAProject(name string) error {
	return a.app.CreateProject(name)
}

func (a *accountFeature) personShouldSeeTheirProject(name string) error {
	projects, err := a.app.GetProjects(name)
	if err != nil {
		return err;
	}
	if len(projects) != 1 {
		return fmt.Errorf("Expected exactly 1 project but got %d", len(projects))
	}
	return nil
}

func (a *accountFeature) personActivatesTheirAccount(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldBeAuthenticated(name string) error {
	if !a.app.IsAuthenticated(name){
		return fmt.Errorf("Expected %s to be authenticated", name)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	af := &accountFeature{
		app: domain.NewDriver(),
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
