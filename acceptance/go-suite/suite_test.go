package features_test

import (
	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
	"github.com/stretchr/testify/suite"
)

type FeatureSuite struct {
	suite.Suite
	driver     driver.TestDriver
	lastErrors map[string]error
}

func (s *FeatureSuite) getLastError(name string) error {
	if s.lastErrors == nil {
		return nil
	}
	return s.lastErrors[name]
}

func (s *FeatureSuite) setLastError(name string, err error) {
	if s.lastErrors == nil {
		s.lastErrors = make(map[string]error)
	}
	s.lastErrors[name] = err
}

// SetupTest is called before each test method
func (s *FeatureSuite) SetupTest() {
	s.lastErrors = make(map[string]error)
	s.driver.ClearAll()
}

// Gherkin keyword methods for chaining
func (s *FeatureSuite) given() *FeatureSuite {
	return s
}

func (s *FeatureSuite) when() *FeatureSuite {
	return s
}

func (s *FeatureSuite) and() *FeatureSuite {
	return s
}

func (s *FeatureSuite) then() *FeatureSuite {
	return s
}

// NewFeatureSuite creates a new test suite instance
func NewFeatureSuite(driver driver.TestDriver) *FeatureSuite {
	s := &FeatureSuite{
		driver: driver,
	}
	return s
}
