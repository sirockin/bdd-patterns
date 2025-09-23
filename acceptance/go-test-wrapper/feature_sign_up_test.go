package features_test

import (
	"testing"
)

func TestSignUp(t *testing.T) {
	withTestContext(t, func(t *testing.T, ctx *testContext) {
		// Given
		personHasCreatedAnAccount(t, ctx, "Sue")

		// When
		personActivatesTheirAccount(t, ctx, "Sue")

		// Then
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
