package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
)


func TestSignUp(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		// Arrange
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

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
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		// Arrange
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given a person has created an account but not activated it
		personHasCreatedAnAccount(t, ctx, "Sue")

		// When they try to sign in
		personTriesToSignIn(t, ctx, "Sue")

		// Then they should not be authenticated
		personShouldNotBeAuthenticated(t, ctx, "Sue")

		// And they should see an error telling them to activate the account
		personShouldSeeAnErrorTellingThemToActivateTheAccount(t, ctx, "Sue")
	})
}

func TestNewPersonCannotSeeProjects(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		// Arrange
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given a person has signed up
		personHasSignedUp(t, ctx, "Sue")

		// When they look at their projects
		// Then they should not see any projects
		personShouldNotSeeAnyProjects(t, ctx, "Sue")
	})
}
