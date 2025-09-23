package features_test

func (s *FeatureSuite) personHasCreatedAnAccount(name string) *FeatureSuite {
	err := s.driver.CreateAccount(name)
	s.Require().NoError(err)
	return s
}

func (s *FeatureSuite) personHasSignedUp(name string) *FeatureSuite {
	err := s.driver.CreateAccount(name)
	s.Require().NoError(err)
	_, err = s.driver.GetAccount(name)
	s.Require().NoError(err)
	err = s.driver.Activate(name)
	s.Require().NoError(err)
	return s
}

func (s *FeatureSuite) personShouldBeAuthenticated(name string) *FeatureSuite {
	actual := s.driver.IsAuthenticated(name)
	s.Assert().True(actual, "person %s should be authenticated", name)
	return s
}

func (s *FeatureSuite) personShouldNotBeAuthenticated(name string) *FeatureSuite {
	actual := s.driver.IsAuthenticated(name)
	s.Assert().False(actual, "person %s should not be authenticated", name)
	return s
}

func (s *FeatureSuite) personShouldNotSeeAnyProjects(name string) *FeatureSuite {
	projects, err := s.driver.GetProjects(name)
	s.Require().NoError(err)
	s.Assert().Empty(projects, "person %s should not see any projects", name)
	return s
}

func (s *FeatureSuite) personShouldSeeTheirProject(name string) *FeatureSuite {
	projects, err := s.driver.GetProjects(name)
	s.Require().NoError(err)
	s.Assert().Len(projects, 1, "person %s should see exactly one project", name)
	return s
}

func (s *FeatureSuite) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) *FeatureSuite {
	lastError := s.getLastError(name)
	expectedText := "you need to activate your account"
	s.Require().NotNil(lastError, "expected error containing text '%s' but there is no error", expectedText)
	s.Assert().Contains(lastError.Error(), expectedText, "expected error text containing '%s'", expectedText)
	return s
}

func (s *FeatureSuite) personTriesToSignIn(name string) *FeatureSuite {
	err := s.driver.Authenticate(name)
	s.setLastError(name, err)
	return s // The step succeeds even if the result is bad to allow the next step to check the error
}

func (s *FeatureSuite) personCreatesAProject(name string) *FeatureSuite {
	err := s.driver.CreateProject(name)
	s.Require().NoError(err)
	return s
}

func (s *FeatureSuite) personActivatesTheirAccount(name string) *FeatureSuite {
	_, err := s.driver.GetAccount(name)
	s.Require().NoError(err)
	err = s.driver.Activate(name)
	s.Require().NoError(err)
	return s
}
