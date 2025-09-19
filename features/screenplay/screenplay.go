// screenplay design pattern utitlity types
package screenplay

import (
	"fmt"
	"strings"

	"github.com/sirockin/cucumber-screenplay-go/features/driver"
)

type Action func(Abilities) error
type Question func(Abilities) (interface{}, error)

type Abilities struct {
	Name      string
	App       driver.ApplicationDriver
	LastError error
}

func (a *Abilities) AttemptsTo(actions ...Action) error {
	for i := 0; i < len(actions); i++ {
		err := actions[i](*a)
		if err != nil {
			a.LastError = err
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
			Name: name,
			App:  app,
		},
	}
	return ret
}

func (a *Actor) AttemptsTo(actions ...Action) error {
	return a.abilities.AttemptsTo(actions...)
}

func (a *Actor) ExpectsAnswer(question Question, expected interface{}) error {
	result, err := question(a.abilities)
	if err != nil {
		return nil
	}
	if result != expected {
		return fmt.Errorf("expected %v to equal %v", result, expected)
	}
	return nil
}

func (a *Actor) ExpectsLastErrorToContain(expectedText string) error {
	if a.abilities.LastError == nil {
		return fmt.Errorf("expected error containing text '%s' but there is no error", expectedText)
	}
	if !strings.Contains(a.abilities.LastError.Error(), expectedText) {
		return fmt.Errorf("expected error text containing '%s' but got %s", expectedText, a.abilities.LastError.Error())
	}
	return nil
}
