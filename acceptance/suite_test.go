package features_test

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
	"github.com/sirockin/cucumber-screenplay-go/acceptance/screenplay"
)

type suite struct {
	actors map[string]*screenplay.Actor
	driver driver.AcceptanceTestDriver
}

func (s *suite) Actor(name string) *screenplay.Actor {
	if s.actors[name] == nil {
		s.actors[name] = screenplay.NewActor(name, s.driver)
	}
	return s.actors[name]
}

func RunSuite(t *testing.T, driver driver.AcceptanceTestDriver, featurePaths []string) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			s := &suite{
				driver: driver,
			}

			ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
				s.actors = make(map[string]*screenplay.Actor)
				s.driver.ClearAll()
				return ctx, nil
			})

			ctx.Step(`^(Bob|Tanya|Sue) has created an account$`, s.personHasCreatedAnAccount)
			ctx.Step(`^(Bob|Tanya|Sue) has signed up$`, s.personHasSignedUp)
			ctx.Step(`^(Bob|Tanya|Sue) should not be authenticated$`, s.personShouldNotBeAuthenticated)
			ctx.Step(`^(Bob|Tanya|Sue) should not see any projects$`, s.personShouldNotSeeAnyProjects)
			ctx.Step(`^(Bob|Tanya|Sue) should see an error telling (him|her|them) to activate the account$`, s.personShouldSeeAnErrorTellingThemToActivateTheAccount)
			ctx.Step(`^(Bob|Tanya|Sue) tries to sign in$`, s.personTriesToSignIn)
			ctx.Step(`^(Bob|Tanya|Sue) creates a project$`, s.personCreatesAProject)
			ctx.Step(`^(Bob|Tanya|Sue) should see (his|her|the) project$`, s.personShouldSeeTheirProject)
			ctx.Step(`^(Bob|Tanya|Sue) activates (his|her) account$`, s.personActivatesTheirAccount)
			ctx.Step(`^(Bob|Tanya|Sue) should be authenticated$`, s.personShouldBeAuthenticated)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    featurePaths,
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
