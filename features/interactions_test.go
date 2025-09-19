package features_test

var CreateAccount = struct {
	forThemselves Action
}{
	forThemselves: func(abilities Abilities) error {
		return abilities.app.CreateAccount(abilities.name)
	},
}

var Activate = struct {
	theirAccount Action
}{
	theirAccount: func(abilities Abilities) error {
		if _, err := abilities.app.GetAccount(abilities.name); err != nil {
			return nil
		}
		return abilities.app.Activate(abilities.name)
	},
}

func signUp(abilities Abilities) error {
	return abilities.attemptsTo(
		CreateAccount.forThemselves,
		Activate.theirAccount,
	)
}

func signIn(abilities Abilities) error {
	return abilities.app.Authenticate(abilities.name)
}

func createProject(abilities Abilities) error {
	return abilities.app.CreateProject(abilities.name)
}
