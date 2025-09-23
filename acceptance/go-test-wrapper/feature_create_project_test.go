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

		// Given Sue has signed up
		assertNoError(t, ctx.personHasSignedUp("Sue"))

		// When Sue creates a project
		assertNoError(t, ctx.personCreatesAProject("Sue"))

		// Then Sue should see the project
		assertNoError(t, ctx.personShouldSeeTheirProject("Sue"))
	})
}

func TestTryToSeeSomeoneElsesProject(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		// Arrange
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given Sue has signed up
		assertNoError(t, ctx.personHasSignedUp("Sue"))

		// And Bob has signed up
		assertNoError(t, ctx.personHasSignedUp("Bob"))

		// When Sue creates a project
		assertNoError(t, ctx.personCreatesAProject("Sue"))

		// Then Bob should not see any projects
		assertNoError(t, ctx.personShouldNotSeeAnyProjects("Bob"))
	})
}