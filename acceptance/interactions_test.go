package features_test

import "github.com/sirockin/cucumber-screenplay-go/acceptance/screenplay"

var CreateAccount = struct {
	forThemselves screenplay.Action
}{
	forThemselves: func(abilities screenplay.Abilities) error {
		return abilities.App.CreateAccount(abilities.Name)
	},
}

var Activate = struct {
	theirAccount screenplay.Action
}{
	theirAccount: func(abilities screenplay.Abilities) error {
		if _, err := abilities.App.GetAccount(abilities.Name); err != nil {
			return nil
		}
		return abilities.App.Activate(abilities.Name)
	},
}

func signUp(abilities screenplay.Abilities) error {
	return abilities.AttemptsTo(
		CreateAccount.forThemselves,
		Activate.theirAccount,
	)
}

func signIn(abilities screenplay.Abilities) error {
	return abilities.App.Authenticate(abilities.Name)
}

func createProject(abilities screenplay.Abilities) error {
	return abilities.App.CreateProject(abilities.Name)
}
