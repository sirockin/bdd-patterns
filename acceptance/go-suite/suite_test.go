package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
)

type suite struct {
	t          *testing.T
	driver     driver.TestDriver
	lastErrors map[string]error
}

func (s *suite) getLastError(name string) error {
	if s.lastErrors == nil {
		return nil
	}
	return s.lastErrors[name]
}

func (s *suite) setLastError(name string, err error) {
	if s.lastErrors == nil {
		s.lastErrors = make(map[string]error)
	}
	s.lastErrors[name] = err
}

// Gherkin keyword methods for chaining
func (s *suite) given() *suite {
	return s
}

func (s *suite) when() *suite {
	return s
}

func (s *suite) and() *suite {
	return s
}

func (s *suite) then() *suite {
	return s
}

// NewSuite creates a new test suite instance
func NewSuite(t *testing.T, driver driver.TestDriver) *suite {
	s := &suite{
		t:      t,
		driver: driver,
	}
	s.lastErrors = make(map[string]error)
	s.driver.ClearAll()
	return s
}
