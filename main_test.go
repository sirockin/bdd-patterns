package main

import (
	"github.com/cucumber/godog"
)

func personHasCreatedAnAccount(name string) error {
	return godog.ErrPending
}

func personHasSignedUp(name string) error {
	return godog.ErrPending
}

func personShouldNotBeAuthenticated(name string) error {
	return godog.ErrPending
}

func personShouldNotSeeAnyProjects(name string) error {
	return godog.ErrPending
}

func personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	return godog.ErrPending
}

func personTriesToSignIn(name string) error {
	return godog.ErrPending
}

func personCreatesAProject(name string) error {
	return godog.ErrPending
}

func personShouldSeeTheirProject(name string) error {
	return godog.ErrPending
}

func personActivatesTheirAccount(name string) error {
	return godog.ErrPending
}

func personShouldBeAuthenticated(name string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^(Bob|Tanya|Sue) has created an account$`, personHasCreatedAnAccount)
	ctx.Step(`^(Bob|Tanya|Sue) has signed up$`, personHasSignedUp)
	ctx.Step(`^(Bob|Tanya|Sue) should not be authenticated$`, personShouldNotBeAuthenticated)
	ctx.Step(`^(Bob|Tanya|Sue) should not see any projects$`, personShouldNotSeeAnyProjects)
	ctx.Step(`^(Bob|Tanya|Sue) should see an error telling (him|her|them) to activate the account$`, personShouldSeeAnErrorTellingThemToActivateTheAccount)
	ctx.Step(`^(Bob|Tanya|Sue) tries to sign in$`, personTriesToSignIn)
	ctx.Step(`^(Bob|Tanya|Sue) creates a project$`, personCreatesAProject)
	ctx.Step(`^(Bob|Tanya|Sue) should see (his|her|the) project$`, personShouldSeeTheirProject)
	ctx.Step(`^(Bob|Tanya|Sue) activates (his|her) account$`, personActivatesTheirAccount)
	ctx.Step(`^(Bob|Tanya|Sue) should be authenticated$`, personShouldBeAuthenticated)
}
