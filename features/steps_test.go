package features_test

func (s *suite) personHasCreatedAnAccount(name string) error {
	return s.Actor(name).attemptsTo(CreateAccount.forThemselves)
}

func (s *suite) personHasSignedUp(name string) error {
	return s.Actor(name).attemptsTo(signUp)
}

func (s *suite) personShouldBeAuthenticated(name string) error {
	return s.Actor(name).expectsAnswer(amIAuthenticated, true)
}

func (s *suite) personShouldNotBeAuthenticated(name string) error {
	return s.Actor(name).expectsAnswer(amIAuthenticated, false)
}

func (s *suite) personShouldNotSeeAnyProjects(name string) error {
	return s.Actor(name).expectsAnswer(howManyProjectsDoIHave, 0)
}

func (s *suite) personShouldSeeTheirProject(name string) error {
	return s.Actor(name).expectsAnswer(howManyProjectsDoIHave, 1)
}

func (s *suite) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	return s.Actor(name).expectsLastErrorToContain("you need to activate your account")
}

func (s *suite) personTriesToSignIn(name string) error {
	_ = s.Actor(name).attemptsTo(signIn)
	return nil // The step succeeds even if the result is bad to allow the next step to check the error
}

func (s *suite) personCreatesAProject(name string) error {
	return s.Actor(name).attemptsTo(createProject)
}

func (s *suite) personActivatesTheirAccount(name string) error {
	return s.Actor(name).attemptsTo(Activate.theirAccount)
}
