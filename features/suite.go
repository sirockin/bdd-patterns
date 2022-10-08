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

func(f *suite) reset(){
	f.actors = make(map[string]*Actor)
	f.driver.ClearAll()
}

func(f *suite) Actor(name string)*Actor{
	if f.actors[name]==nil {
		f.actors[name]=NewActor(name, f.driver)
	}
	return f.actors[name]	
}

func Test(t *testing.T, driver ApplicationDriver, featurePaths []string) {
	// TODO: Try to get rid of need to pass in path to features
	suite := godog.TestSuite{
	  ScenarioInitializer: func (ctx *godog.ScenarioContext) {
		f := &suite{
			driver: driver,
		}
	
		ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {		
			f.reset()
			return ctx, nil
		})
	
		ctx.Step(`^(Bob|Tanya|Sue) has created an account$`, f.personHasCreatedAnAccount)
		ctx.Step(`^(Bob|Tanya|Sue) has signed up$`, f.personHasSignedUp)
		ctx.Step(`^(Bob|Tanya|Sue) should not be authenticated$`, f.personShouldNotBeAuthenticated)
		ctx.Step(`^(Bob|Tanya|Sue) should not see any projects$`, f.personShouldNotSeeAnyProjects)
		ctx.Step(`^(Bob|Tanya|Sue) should see an error telling (him|her|them) to activate the account$`, f.personShouldSeeAnErrorTellingThemToActivateTheAccount)
		ctx.Step(`^(Bob|Tanya|Sue) tries to sign in$`, f.personTriesToSignIn)
		ctx.Step(`^(Bob|Tanya|Sue) creates a project$`, f.personCreatesAProject)
		ctx.Step(`^(Bob|Tanya|Sue) should see (his|her|the) project$`, f.personShouldSeeTheirProject)
		ctx.Step(`^(Bob|Tanya|Sue) activates (his|her) account$`, f.personActivatesTheirAccount)
		ctx.Step(`^(Bob|Tanya|Sue) should be authenticated$`, f.personShouldBeAuthenticated)
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
