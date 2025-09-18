package features

import (
	"fmt"
	"strings"

	"github.com/sirockin/cucumber-screenplay-go/features/driver"
)

type Action func(Abilities) error
type Question func(Abilities) (interface{}, error)

type Abilities struct {
	name      string
	app       driver.ApplicationDriver
	lastError error
}

func (a *Abilities) attemptsTo(actions ...Action) error {
	for i := 0; i < len(actions); i++ {
		err := actions[i](*a)
		if err != nil {
			a.lastError = err
			return err
		}
	}
	return nil
}

type Actor struct {
	abilities Abilities
}

func NewActor(name string, app driver.ApplicationDriver) *Actor {
	ret := &Actor{
		abilities: Abilities{
			name: name,
			app:  app,
		},
	}
	return ret
}

func (a *Actor) attemptsTo(actions ...Action) error {
	return a.abilities.attemptsTo(actions...)
}

func (a *Actor) expectsAnswer(question Question, expected interface{}) error {
	result, err := question(a.abilities)
	if err != nil {
		return nil
	}
	if result != expected {
		return fmt.Errorf("expected %v to equal %v", result, expected)
	}
	return nil
}

func (a *Actor) expectsLastErrorToContain(expectedText string) error {
	if a.abilities.lastError == nil {
		return fmt.Errorf("expected error containing text '%s' but there is no error", expectedText)
	}
	if !strings.Contains(a.abilities.lastError.Error(), expectedText) {
		return fmt.Errorf("expected error text containing '%s' but got %s", expectedText, a.abilities.lastError.Error())
	}
	return nil

}
