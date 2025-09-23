package features_test

func (s *suite) personHasCreatedAnAccount(name string) error {
	return s.Actor(name).AttemptsTo(CreateAccount.forThemselves)
}

func (s *suite) personHasSignedUp(name string) error {
	return s.Actor(name).AttemptsTo(signUp)
}

func (s *suite) personShouldBeAuthenticated(name string) error {
	return s.Actor(name).ExpectsAnswer(amIAuthenticated, true)
}

func (s *suite) personShouldNotBeAuthenticated(name string) error {
	return s.Actor(name).ExpectsAnswer(amIAuthenticated, false)
}

func (s *suite) personShouldNotSeeAnyProjects(name string) error {
	return s.Actor(name).ExpectsAnswer(howManyProjectsDoIHave, 0)
}

func (s *suite) personShouldSeeTheirProject(name string) error {
	return s.Actor(name).ExpectsAnswer(howManyProjectsDoIHave, 1)
}

func (s *suite) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	return s.Actor(name).ExpectsLastErrorToContain("you need to activate your account")
}

func (s *suite) personTriesToSignIn(name string) error {
	_ = s.Actor(name).AttemptsTo(signIn)
	return nil // The step succeeds even if the result is bad to allow the next step to check the error
}

func (s *suite) personCreatesAProject(name string) error {
	return s.Actor(name).AttemptsTo(createProject)
}

func (s *suite) personActivatesTheirAccount(name string) error {
	return s.Actor(name).AttemptsTo(Activate.theirAccount)
}
