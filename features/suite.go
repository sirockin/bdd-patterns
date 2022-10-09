package features

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
)

type suite struct {
	actors map[string]*Actor
	driver ApplicationDriver
}

func(s *suite) reset(){
	s.actors = make(map[string]*Actor)
	s.driver.ClearAll()
}

func(s *suite) Actor(name string)*Actor{
	if s.actors[name]==nil {
		s.actors[name]=NewActor(name, s.driver)
	}
	return s.actors[name]	
}

func Test(t *testing.T, driver ApplicationDriver, featurePaths []string) {
	// TODO: Try to get rid of need to pass in path to features
	suite := godog.TestSuite{
	  ScenarioInitializer: func (ctx *godog.ScenarioContext) {
		s := &suite{
			driver: driver,
		}
	
		ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {		
			s.reset()
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
