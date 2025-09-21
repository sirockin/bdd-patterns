package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"
)

func newSuite(t *testing.T, driver driver.TestDriver) *suite {
	s := &suite{
		t:      t,
		driver: driver,
	}
	s.lastErrors = make(map[string]error)
	s.driver.ClearAll()
	return s
}

func TestCreateProjectFeature_CreateOneProject(t *testing.T) {
	driver := testhelpers.NewDomainTestDriver()
	s := newSuite(t, driver)

	s.given().
		sueHasSignedUp().
		when().
		sueCreatesAProject().
		then().
		sueShouldSeeTheProject()
}

func TestCreateProjectFeature_TryToSeeSomeoneElsesProject(t *testing.T) {
	driver := testhelpers.NewDomainTestDriver()
	s := newSuite(t, driver)

	s.given().
		sueHasSignedUp().
		and().
		bobHasSignedUp().
		when().
		sueCreatesAProject().
		then().
		bobShouldNotSeeAnyProjects()
}