package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/acceptance/driver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func personHasCreatedAnAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	err := ctx.driver.CreateAccount(name)
	require.NoError(t, err)
}

func personHasSignedUp(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	err := ctx.driver.CreateAccount(name)
	require.NoError(t, err)
	_, err = ctx.driver.GetAccount(name)
	require.NoError(t, err)
	err = ctx.driver.Activate(name)
	require.NoError(t, err)
}

func personShouldBeAuthenticated(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	actual := ctx.driver.IsAuthenticated(name)
	assert.True(t, actual, "person %s should be authenticated", name)
}

func personShouldNotBeAuthenticated(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	actual := ctx.driver.IsAuthenticated(name)
	assert.False(t, actual, "person %s should not be authenticated", name)
}

func personShouldNotSeeAnyProjects(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	projects, err := ctx.driver.GetProjects(name)
	require.NoError(t, err)
	assert.Empty(t, projects, "person %s should not see any projects", name)
}

func personShouldSeeTheirProject(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	projects, err := ctx.driver.GetProjects(name)
	require.NoError(t, err)
	assert.Len(t, projects, 1, "person %s should see exactly one project", name)
}

func personShouldSeeAnErrorTellingThemToActivateTheAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	lastError := ctx.getLastError(name)
	expectedText := "you need to activate your account"
	require.NotNil(t, lastError, "expected error containing text '%s' but there is no error", expectedText)
	assert.Contains(t, lastError.Error(), expectedText, "expected error text containing '%s'", expectedText)
}

func personTriesToSignIn(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	err := ctx.driver.Authenticate(name)
	ctx.setLastError(name, err)
	// The step succeeds even if the result is bad to allow the next step to check the error
}

func personCreatesAProject(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	err := ctx.driver.CreateProject(name)
	require.NoError(t, err)
}

func personActivatesTheirAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	_, err := ctx.driver.GetAccount(name)
	require.NoError(t, err)
	err = ctx.driver.Activate(name)
	require.NoError(t, err)
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

