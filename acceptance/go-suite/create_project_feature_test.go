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
		personHasSignedUp("Sue").
		when().
		personCreatesAProject("Sue").
		then().
		personShouldSeeTheirProject("Sue")
}

func TestCreateProjectFeature_TryToSeeSomeoneElsesProject(t *testing.T) {
	driver := testhelpers.NewDomainTestDriver()
	s := newSuite(t, driver)

	s.given().
		personHasSignedUp("Sue").
		and().
		personHasSignedUp("Bob").
		when().
		personCreatesAProject("Sue").
		then().
		personShouldNotSeeAnyProjects("Bob")
}