package features


////////
// Steps

func (f *suite) personHasCreatedAnAccount(name string) error {
	return f.Actor(name).attemptsTo(CreateAccount.forThemselves)
}

func (f *suite) personHasSignedUp(name string) error {
	return f.Actor(name).attemptsTo(signUp)
}

func (f *suite) personShouldBeAuthenticated(name string) error {
	return f.Actor(name).expectsAnswer(amIAuthenticated,true)
}

func (f *suite) personShouldNotBeAuthenticated(name string) error {
	return f.Actor(name).expectsAnswer(amIAuthenticated,false)
}

func (f *suite) personShouldNotSeeAnyProjects(name string) error {
	return f.Actor(name).expectsAnswer(howManyProjectsDoIHave,0)
}

func (f *suite) personShouldSeeTheirProject(name string) error {
	return f.Actor(name).expectsAnswer(howManyProjectsDoIHave,1)
}

func (f *suite) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	return f.Actor(name).expectsLastErrorToContain("you need to activate your account")
}

func (f *suite) personTriesToSignIn(name string) error {
	f.Actor(name).attemptsTo(signIn)
	return nil;	// The step succeeds even if the result is bad to allow the next step to check the error
}

func (f *suite) personCreatesAProject(name string) error {
	return f.Actor(name).attemptsTo(createProject)
}

func (f *suite) personActivatesTheirAccount(name string) error {
	return f.Actor(name).attemptsTo(Activate.theirAccount)
}



