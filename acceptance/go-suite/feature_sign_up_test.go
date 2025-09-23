package features_test

// TestSuccessfulSignUp tests successful account activation and authentication
func (s *FeatureSuite) TestSuccessfulSignUp() {
	s.
		given().personHasCreatedAnAccount("Tanya").
		when().personActivatesTheirAccount("Tanya").
		then().personShouldBeAuthenticated("Tanya")
}

// TestTryToSignInWithoutActivatingAccount tests sign-in failure without account activation
func (s *FeatureSuite) TestTryToSignInWithoutActivatingAccount() {
	s.
		given().personHasCreatedAnAccount("Bob").
		when().personTriesToSignIn("Bob").
		then().personShouldNotBeAuthenticated("Bob").
		and().personShouldSeeAnErrorTellingThemToActivateTheAccount("Bob")
}
