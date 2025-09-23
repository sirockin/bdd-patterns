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
		// Given
		personHasSignedUp(t, ctx, "Sue")
		personHasSignedUp(t, ctx, "Bob")

		// When
		personCreatesAProject(t, ctx, "Sue")

		// Then
		personShouldNotSeeAnyProjects(t, ctx, "Bob")
	})
}
