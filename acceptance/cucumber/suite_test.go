package features_test

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
)

type suite struct {
	driver      driver.TestDriver
	lastErrors  map[string]error
}

func (s *suite) getLastError(name string) error {
	if s.lastErrors == nil {
		return nil
	}
	return s.lastErrors[name]
}

func (s *suite) setLastError(name string, err error) {
	if s.lastErrors == nil {
		s.lastErrors = make(map[string]error)
	}
	s.lastErrors[name] = err
}

func RunSuite(t *testing.T, driver driver.TestDriver, featurePaths []string) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			s := &suite{
				driver: driver,
			}

			ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
				s.lastErrors = make(map[string]error)
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
