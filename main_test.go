package main

import (
	"github.com/cucumber/godog"
)

func bobHasCreatedAnAccount() error {
	return godog.ErrPending
}

func bobHasSignedUp() error {
	return godog.ErrPending
}

func bobShouldNotBeAuthenticated() error {
	return godog.ErrPending
}

func bobShouldNotSeeAnyProjects() error {
	return godog.ErrPending
}

func bobShouldSeeAnErrorTellingHimToActivateTheAccount() error {
	return godog.ErrPending
}

func bobTriesToSignIn() error {
	return godog.ErrPending
}

func sueCreatesAProject() error {
	return godog.ErrPending
}

func sueHasSignedUp() error {
	return godog.ErrPending
}

func sueShouldSeeTheProject() error {
	return godog.ErrPending
}

func tanyaActivatesHerAccount() error {
	return godog.ErrPending
}

func tanyaHasCreatedAnAccount() error {
	return godog.ErrPending
}

func tanyaShouldBeAuthenticated() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Bob has created an account$`, bobHasCreatedAnAccount)
	ctx.Step(`^Bob has signed up$`, bobHasSignedUp)
	ctx.Step(`^Bob should not be authenticated$`, bobShouldNotBeAuthenticated)
	ctx.Step(`^Bob should not see any projects$`, bobShouldNotSeeAnyProjects)
	ctx.Step(`^Bob should see an error telling him to activate the account$`, bobShouldSeeAnErrorTellingHimToActivateTheAccount)
	ctx.Step(`^Bob tries to sign in$`, bobTriesToSignIn)
	ctx.Step(`^Sue creates a project$`, sueCreatesAProject)
	ctx.Step(`^Sue has signed up$`, sueHasSignedUp)
	ctx.Step(`^Sue should see the project$`, sueShouldSeeTheProject)
	ctx.Step(`^Tanya activates her account$`, tanyaActivatesHerAccount)
	ctx.Step(`^Tanya has created an account$`, tanyaHasCreatedAnAccount)
	ctx.Step(`^Tanya should be authenticated$`, tanyaShouldBeAuthenticated)
}
