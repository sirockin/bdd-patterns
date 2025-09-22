package features_test

import (
	"strings"
)


func (s *suite) personHasCreatedAnAccount(name string) *suite {
	if err := s.driver.CreateAccount(name); err != nil {
		s.t.Fatal(err)
	}
	return s
}

func (s *suite) personHasSignedUp(name string) *suite {
	if err := s.driver.CreateAccount(name); err != nil {
		s.t.Fatal(err)
	}
	if _, err := s.driver.GetAccount(name); err != nil {
		s.t.Fatal(err)
	}
	if err := s.driver.Activate(name); err != nil {
		s.t.Fatal(err)
	}
	return s
}

func (s *suite) personShouldBeAuthenticated(name string) *suite {
	expected := true
	actual := s.driver.IsAuthenticated(name)
	if actual != expected {
		s.t.Fatalf("expected %v to equal %v", actual, expected)
	}
	return s
}

func (s *suite) personShouldNotBeAuthenticated(name string) *suite {
	expected := false
	actual := s.driver.IsAuthenticated(name)
	if actual != expected {
		s.t.Fatalf("expected %v to equal %v", actual, expected)
	}
	return s
}

func (s *suite) personShouldNotSeeAnyProjects(name string) *suite {
	projects, err := s.driver.GetProjects(name)
	if err != nil {
		s.t.Fatal(err)
	}
	expected := 0
	actual := len(projects)
	if actual != expected {
		s.t.Fatalf("expected %v to equal %v", actual, expected)
	}
	return s
}

func (s *suite) personShouldSeeTheirProject(name string) *suite {
	projects, err := s.driver.GetProjects(name)
	if err != nil {
		s.t.Fatal(err)
	}
	expected := 1
	actual := len(projects)
	if actual != expected {
		s.t.Fatalf("expected %v to equal %v", actual, expected)
	}
	return s
}

func (s *suite) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) *suite {
	lastError := s.getLastError(name)
	expectedText := "you need to activate your account"
	if lastError == nil {
		s.t.Fatalf("expected error containing text '%s' but there is no error", expectedText)
	}
	if !strings.Contains(lastError.Error(), expectedText) {
		s.t.Fatalf("expected error text containing '%s' but got %s", expectedText, lastError.Error())
	}
	return s
}

func (s *suite) personTriesToSignIn(name string) *suite {
	err := s.driver.Authenticate(name)
	s.setLastError(name, err)
	return s // The step succeeds even if the result is bad to allow the next step to check the error
}

func (s *suite) personCreatesAProject(name string) *suite {
	if err := s.driver.CreateProject(name); err != nil {
		s.t.Fatal(err)
	}
	return s
}

func (s *suite) personActivatesTheirAccount(name string) *suite {
	if _, err := s.driver.GetAccount(name); err != nil {
		s.t.Fatal(err)
	}
	if err := s.driver.Activate(name); err != nil {
		s.t.Fatal(err)
	}
	return s
}
