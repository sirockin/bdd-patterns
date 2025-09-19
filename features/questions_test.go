package features_test

func amIAuthenticated(abilities Abilities) (interface{}, error) {
	return abilities.app.IsAuthenticated(abilities.name), nil
}

func howManyProjectsDoIHave(abilities Abilities) (interface{}, error) {
	projects, err := abilities.app.GetProjects(abilities.name)
	if err != nil {
		return 0, err
	}
	return len(projects), nil
}
