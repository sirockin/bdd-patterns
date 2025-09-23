package features_test

import (
	"testing"
)

func TestCreateOneProject(t *testing.T) {
	withTestContext(t, func(t *testing.T, ctx *testContext) {
		// Given
		personHasSignedUp(t, ctx, "Sue")

		// When
		personCreatesAProject(t, ctx, "Sue")

		// Then
		personShouldSeeTheirProject(t, ctx, "Sue")
	})
}

func TestTryToSeeSomeoneElsesProject(t *testing.T) {
	withTestContext(t, func(t *testing.T, ctx *testContext) {
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