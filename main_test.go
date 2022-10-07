package main

import (
	"context"

	"github.com/cucumber/godog"
)

type accountFeature struct {
}

func (a *accountFeature) personHasCreatedAnAccount(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personHasSignedUp(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldNotBeAuthenticated(name string) error {
	return godog.ErrPending
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
	return godog.ErrPending
}

func (a *accountFeature) personShouldSeeTheirProject(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personActivatesTheirAccount(name string) error {
	return godog.ErrPending
}

func (a *accountFeature) personShouldBeAuthenticated(name string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	af := &accountFeature{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
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
