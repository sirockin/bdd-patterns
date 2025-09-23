package features_test

// TestCreateOneProject tests creating a single project
func (s *FeatureSuite) TestCreateOneProject() {
	s.
		given().personHasSignedUp("Sue").
		when().personCreatesAProject("Sue").
		then().personShouldSeeTheirProject("Sue")
}

// TestTryToSeeSomeoneElsesProject tests that users cannot see other users' projects
func (s *FeatureSuite) TestTryToSeeSomeoneElsesProject() {
	s.
		given().personHasSignedUp("Sue").
		and().personHasSignedUp("Bob").
		when().personCreatesAProject("Sue").
		then().personShouldNotSeeAnyProjects("Bob")
}
