package main_test

import (
	"fmt"
	"strings"
)

// Screenplay objects
type Abilities struct {
	name string
	app Application
	lastError error
}

func (a* Abilities)attemptsTo(actions ...Action)error{
	for i:=0; i<len(actions); i++{
		err := actions[i](*a)
		if err != nil {
			a.lastError=err
			return err
		}
	}
	return nil
}

type Actor struct {
	abilities Abilities
}

func(a* Actor)attemptsTo(actions...Action)error{
	return a.abilities.attemptsTo(actions...)
}

func (a* Actor) expectsAnswer(question Question, expected interface{})error{
	result, err := question(a.abilities)
	if err != nil {
		return nil
	}
	if result != expected {
		return fmt.Errorf("Expected %v to equal %v", result, expected)
	}
	return nil
}

func( a* Actor) expectsLastErrorToContain(expectedText string)error{
	if a.abilities.lastError == nil {
		return fmt.Errorf("Expected error containing text '%s' but there is no error", expectedText)
	}
	if !strings.Contains(a.abilities.lastError.Error(), expectedText){
		return fmt.Errorf("Expected error text containing '%s' but got %s", expectedText, a.abilities.lastError.Error())
	}
	return nil

}

type Action func(Abilities)error
type Question func(Abilities)(interface{}, error)


func NewActor(name string, app Application)*Actor{
	ret := &Actor{
		abilities:Abilities{
			name: name, 
			app: app,	
		},
	}
	return ret
}

