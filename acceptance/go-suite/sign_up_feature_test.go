package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/testhelpers"
)

func TestSignUpFeature_SuccessfulSignUp(t *testing.T) {
	driver := testhelpers.NewDomainTestDriver()
	s := newSuite(t, driver)

	s.given().
		tanyaHasCreatedAnAccount().
		when().
		tanyaActivatesHerAccount().
		then().
		tanyaShouldBeAuthenticated()
}

func TestSignUpFeature_TryToSignInWithoutActivatingAccount(t *testing.T) {
	driver := testhelpers.NewDomainTestDriver()
	s := newSuite(t, driver)

	s.given().
		bobHasCreatedAnAccount().
		when().
		bobTriesToSignIn().
		then().
		bobShouldNotBeAuthenticated().
		and().
		bobShouldSeeAnErrorTellingHimToActivateTheAccount()
}