package features_test

import "github.com/sirockin/cucumber-screenplay-go/features/screenplay"

func amIAuthenticated(abilities screenplay.Abilities) (interface{}, error) {
	return abilities.App.IsAuthenticated(abilities.Name), nil
}

func howManyProjectsDoIHave(abilities screenplay.Abilities) (interface{}, error) {
	projects, err := abilities.App.GetProjects(abilities.Name)
	if err != nil {
		return 0, err
	}
	return len(projects), nil
}
