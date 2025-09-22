package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"
)

func TestSignUpFeature_SuccessfulSignUp(t *testing.T) {
	driver := testhelpers.NewDomainTestDriver()
	s := newSuite(t, driver)

	s.given().
		personHasCreatedAnAccount("Tanya").
		when().
		personActivatesTheirAccount("Tanya").
		then().
		personShouldBeAuthenticated("Tanya")
}

func TestSignUpFeature_TryToSignInWithoutActivatingAccount(t *testing.T) {
	driver := testhelpers.NewDomainTestDriver()
	s := newSuite(t, driver)

	s.given().
		personHasCreatedAnAccount("Bob").
		when().
		personTriesToSignIn("Bob").
		then().
		personShouldNotBeAuthenticated("Bob").
		and().
		personShouldSeeAnErrorTellingThemToActivateTheAccount("Bob")
}