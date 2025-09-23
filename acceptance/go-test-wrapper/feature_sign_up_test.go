package features_test

import (
	"testing"
)

func TestSignUp(t *testing.T) {
	withTestContext(t, func(t *testing.T, ctx *testContext) {
		// Given a person has created an account
		personHasCreatedAnAccount(t, ctx, "Sue")

		// When they activate their account
		personActivatesTheirAccount(t, ctx, "Sue")

		// Then they should be able to sign in successfully
		personTriesToSignIn(t, ctx, "Sue")
		personShouldBeAuthenticated(t, ctx, "Sue")
	})
}

func TestSignInBeforeActivation(t *testing.T) {
	withTestContext(t, func(t *testing.T, ctx *testContext) {
		// Given
		personHasCreatedAnAccount(t, ctx, "Sue")

		// When
		personTriesToSignIn(t, ctx, "Sue")

		// Then
		personShouldNotBeAuthenticated(t, ctx, "Sue")
		personShouldSeeAnErrorTellingThemToActivateTheAccount(t, ctx, "Sue")
	})
}

func TestNewPersonCannotSeeProjects(t *testing.T) {
	withTestContext(t, func(t *testing.T, ctx *testContext) {
		// Given
		personHasSignedUp(t, ctx, "Sue")

		// Then
		personShouldNotSeeAnyProjects(t, ctx, "Sue")
	})
}
