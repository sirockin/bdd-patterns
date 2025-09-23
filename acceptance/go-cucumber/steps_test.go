package features_test

import (
	"fmt"
	"strings"
)

func (s *suite) personHasCreatedAnAccount(name string) error {
	return s.driver.CreateAccount(name)
}

func (s *suite) personHasSignedUp(name string) error {
	if err := s.driver.CreateAccount(name); err != nil {
		return err
	}
	if _, err := s.driver.GetAccount(name); err != nil {
		return nil
	}
	return s.driver.Activate(name)
}

func (s *suite) personShouldBeAuthenticated(name string) error {
	expected := true
	actual := s.driver.IsAuthenticated(name)
	if actual != expected {
		return fmt.Errorf("expected %v to equal %v", actual, expected)
	}
	return nil
}

func (s *suite) personShouldNotBeAuthenticated(name string) error {
	expected := false
	actual := s.driver.IsAuthenticated(name)
	if actual != expected {
		return fmt.Errorf("expected %v to equal %v", actual, expected)
	}
	return nil
}

func (s *suite) personShouldNotSeeAnyProjects(name string) error {
	projects, err := s.driver.GetProjects(name)
	if err != nil {
		return err
	}
	expected := 0
	actual := len(projects)
	if actual != expected {
		return fmt.Errorf("expected %v to equal %v", actual, expected)
	}
	return nil
}

func (s *suite) personShouldSeeTheirProject(name string) error {
	projects, err := s.driver.GetProjects(name)
	if err != nil {
		return err
	}
	expected := 1
	actual := len(projects)
	if actual != expected {
		return fmt.Errorf("expected %v to equal %v", actual, expected)
	}
	return nil
}

func (s *suite) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	lastError := s.getLastError(name)
	expectedText := "you need to activate your account"
	if lastError == nil {
		return fmt.Errorf("expected error containing text '%s' but there is no error", expectedText)
	}
	if !strings.Contains(lastError.Error(), expectedText) {
		return fmt.Errorf("expected error text containing '%s' but got %s", expectedText, lastError.Error())
	}
	return nil
}

func (s *suite) personTriesToSignIn(name string) error {
	err := s.driver.Authenticate(name)
	s.setLastError(name, err)
	return nil // The step succeeds even if the result is bad to allow the next step to check the error
}

func (s *suite) personCreatesAProject(name string) error {
	return s.driver.CreateProject(name)
}

func (s *suite) personActivatesTheirAccount(name string) error {
	if _, err := s.driver.GetAccount(name); err != nil {
		return nil
	}
	return s.driver.Activate(name)
}
