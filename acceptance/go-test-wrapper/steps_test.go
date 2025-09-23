package features_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
)

type testContext struct {
	driver     driver.TestDriver
	lastErrors map[string]error
}

func newTestContext(testDriver driver.TestDriver) *testContext {
	return &testContext{
		driver:     testDriver,
		lastErrors: make(map[string]error),
	}
}

func (ctx *testContext) personHasCreatedAnAccount(name string) error {
	return ctx.driver.CreateAccount(name)
}

func (ctx *testContext) personHasSignedUp(name string) error {
	if err := ctx.driver.CreateAccount(name); err != nil {
		return err
	}
	if _, err := ctx.driver.GetAccount(name); err != nil {
		return nil
	}
	return ctx.driver.Activate(name)
}

func (ctx *testContext) personShouldBeAuthenticated(name string) error {
	expected := true
	actual := ctx.driver.IsAuthenticated(name)
	if actual != expected {
		return fmt.Errorf("expected %v to equal %v", actual, expected)
	}
	return nil
}

func (ctx *testContext) personShouldNotBeAuthenticated(name string) error {
	expected := false
	actual := ctx.driver.IsAuthenticated(name)
	if actual != expected {
		return fmt.Errorf("expected %v to equal %v", actual, expected)
	}
	return nil
}

func (ctx *testContext) personShouldNotSeeAnyProjects(name string) error {
	projects, err := ctx.driver.GetProjects(name)
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

func (ctx *testContext) personShouldSeeTheirProject(name string) error {
	projects, err := ctx.driver.GetProjects(name)
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

func (ctx *testContext) personShouldSeeAnErrorTellingThemToActivateTheAccount(name string) error {
	lastError := ctx.getLastError(name)
	expectedText := "you need to activate your account"
	if lastError == nil {
		return fmt.Errorf("expected error containing text '%s' but there is no error", expectedText)
	}
	if !strings.Contains(lastError.Error(), expectedText) {
		return fmt.Errorf("expected error text containing '%s' but got %s", expectedText, lastError.Error())
	}
	return nil
}

func (ctx *testContext) personTriesToSignIn(name string) error {
	err := ctx.driver.Authenticate(name)
	ctx.setLastError(name, err)
	return nil
}

func (ctx *testContext) personCreatesAProject(name string) error {
	return ctx.driver.CreateProject(name)
}

func (ctx *testContext) personActivatesTheirAccount(name string) error {
	if _, err := ctx.driver.GetAccount(name); err != nil {
		return nil
	}
	return ctx.driver.Activate(name)
}

func (ctx *testContext) getLastError(name string) error {
	return ctx.lastErrors[name]
}

func (ctx *testContext) setLastError(name string, err error) {
	ctx.lastErrors[name] = err
}

func (ctx *testContext) clearAll() {
	ctx.driver.ClearAll()
	ctx.lastErrors = make(map[string]error)
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
