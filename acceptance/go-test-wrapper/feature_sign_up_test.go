package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
)

func TestCreateOneProject(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		// Arrange
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given
		assertNoError(t, ctx.personHasSignedUp("Sue"))

		// When
		assertNoError(t, ctx.personCreatesAProject("Sue"))

		// Then
		assertNoError(t, ctx.personShouldSeeTheirProject("Sue"))
	})
}

func TestSignUp(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		// Arrange
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given a person has created an account
		assertNoError(t, ctx.personHasCreatedAnAccount("Sue"))

		// When they activate their account
		assertNoError(t, ctx.personActivatesTheirAccount("Sue"))

		// Then they should be able to sign in successfully
		assertNoError(t, ctx.personTriesToSignIn("Sue"))
		assertNoError(t, ctx.personShouldBeAuthenticated("Sue"))
	})
}

func TestSignInBeforeActivation(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		// Arrange
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given a person has created an account but not activated it
		assertNoError(t, ctx.personHasCreatedAnAccount("Sue"))

		// When they try to sign in
		assertNoError(t, ctx.personTriesToSignIn("Sue"))

		// Then they should not be authenticated
		assertNoError(t, ctx.personShouldNotBeAuthenticated("Sue"))

		// And they should see an error telling them to activate the account
		assertNoError(t, ctx.personShouldSeeAnErrorTellingThemToActivateTheAccount("Sue"))
	})
}

func TestNewPersonCannotSeeProjects(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		// Arrange
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given a person has signed up
		assertNoError(t, ctx.personHasSignedUp("Sue"))

		// When they look at their projects
		// Then they should not see any projects
		assertNoError(t, ctx.personShouldNotSeeAnyProjects("Sue"))
	})
}
