package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
)

func TestCreateOneProject(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given
		personHasSignedUp(t, ctx, "Sue")

		// When
		personCreatesAProject(t, ctx, "Sue")

		// Then 
		personShouldSeeTheirProject(t, ctx, "Sue")
	})
}

func TestTryToSeeSomeoneElsesProject(t *testing.T) {
	withTestDriver(t, func(t *testing.T, testDriver driver.TestDriver) {
		ctx := newTestContext(testDriver)
		defer ctx.clearAll()

		// Given Sue has signed up
		personHasSignedUp(t, ctx, "Sue")

		// And Bob has signed up
		personHasSignedUp(t, ctx, "Bob")

		// When Sue creates a project
		personCreatesAProject(t, ctx, "Sue")

		// Then Bob should not see any projects
		personShouldNotSeeAnyProjects(t, ctx, "Bob")
	})
}