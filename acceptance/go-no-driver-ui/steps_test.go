package features_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func personHasCreatedAnAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	// Navigate to the account creation page
	_, err := ctx.page.Goto(ctx.frontendURL + "/signup")
	require.NoError(t, err, "failed to navigate to signup page")

	// Wait for page to load
	_, err = ctx.page.WaitForSelector("input[name='name']", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "signup form not found")

	// Fill in the name field
	err = ctx.page.Fill("input[name='name']", name)
	require.NoError(t, err, "failed to fill name field")

	// Click create account button
	err = ctx.page.Click("button[type='submit']")
	require.NoError(t, err, "failed to click create account button")

	// Wait for success message or redirect
	_, err = ctx.page.WaitForSelector(".success, .message", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "account creation failed or timed out")
}

func personHasSignedUp(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	personHasCreatedAnAccount(t, ctx, name)
	getAccount(t, ctx, name)
	personActivatesTheirAccount(t, ctx, name)
}

func getAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	// Navigate to account details page
	_, err := ctx.page.Goto(ctx.frontendURL + "/account/" + name)
	require.NoError(t, err, "failed to navigate to account page")

	// Wait for account data to load
	_, err = ctx.page.WaitForSelector(".account-info", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "account not found: %s", name)
}

func personShouldBeAuthenticated(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	// Navigate to account page and check authentication status
	_, err := ctx.page.Goto(ctx.frontendURL + "/account/" + name)
	require.NoError(t, err)

	// Check if authenticated indicator is visible
	authenticated, err := ctx.page.IsVisible(".status-authenticated")
	require.NoError(t, err)

	assert.True(t, authenticated, "person %s should be authenticated", name)
}

func personShouldNotBeAuthenticated(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	// Navigate to account page and check authentication status
	_, err := ctx.page.Goto(ctx.frontendURL + "/account/" + name)
	require.NoError(t, err)

	// Check if authenticated indicator is visible
	authenticated, err := ctx.page.IsVisible(".status-authenticated")
	if err != nil {
		authenticated = false
	}

	assert.False(t, authenticated, "person %s should not be authenticated", name)
}

func personShouldNotSeeAnyProjects(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	// Navigate to projects page
	_, err := ctx.page.Goto(ctx.frontendURL + "/account/" + name + "/projects")
	require.NoError(t, err, "failed to navigate to projects page")

	// Wait for projects list
	_, err = ctx.page.WaitForSelector(".projects-list", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "projects list not found")

	// Count project items
	projectElements, err := ctx.page.QuerySelectorAll(".project-item")
	require.NoError(t, err, "failed to find project items")

	assert.Empty(t, projectElements, "person %s should not see any projects", name)
}

func personShouldSeeTheirProject(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	// Navigate to projects page
	_, err := ctx.page.Goto(ctx.frontendURL + "/account/" + name + "/projects")
	require.NoError(t, err, "failed to navigate to projects page")

	// Wait for projects list
	_, err = ctx.page.WaitForSelector(".projects-list", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "projects list not found")

	// Count project items
	projectElements, err := ctx.page.QuerySelectorAll(".project-item")
	require.NoError(t, err, "failed to find project items")

	assert.Len(t, projectElements, 1, "person %s should see exactly one project", name)
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

	// Navigate to login page
	_, err := ctx.page.Goto(ctx.frontendURL + "/login")
	if err != nil {
		ctx.setLastError(name, fmt.Errorf("failed to navigate to login page: %w", err))
		return
	}

	// Wait for login form
	_, err = ctx.page.WaitForSelector("input[name='name']", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		ctx.setLastError(name, fmt.Errorf("login form not found: %w", err))
		return
	}

	// Fill in the name field
	err = ctx.page.Fill("input[name='name']", name)
	if err != nil {
		ctx.setLastError(name, fmt.Errorf("failed to fill name field: %w", err))
		return
	}

	// Click login button
	err = ctx.page.Click("button[type='submit']")
	if err != nil {
		ctx.setLastError(name, fmt.Errorf("failed to click login button: %w", err))
		return
	}

	// Check for error message
	time.Sleep(1 * time.Second) // Give time for response
	errorVisible, _ := ctx.page.IsVisible(".error")
	if errorVisible {
		errorText, _ := ctx.page.TextContent(".error")
		ctx.setLastError(name, fmt.Errorf("%s", errorText))
		return
	}

	ctx.setLastError(name, nil)
}

func personCreatesAProject(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	// Navigate to projects page
	_, err := ctx.page.Goto(ctx.frontendURL + "/account/" + name + "/projects")
	require.NoError(t, err, "failed to navigate to projects page")

	// Wait for create project button
	_, err = ctx.page.WaitForSelector("button.create-project", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "create project button not found")

	// Click create project button
	err = ctx.page.Click("button.create-project")
	require.NoError(t, err, "failed to click create project button")

	// Wait for success message
	_, err = ctx.page.WaitForSelector(".project-created", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "project creation failed or timed out")
}

func personActivatesTheirAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	getAccount(t, ctx, name)

	// Navigate to activation page
	_, err := ctx.page.Goto(ctx.frontendURL + "/activate/" + name)
	require.NoError(t, err, "failed to navigate to activation page")

	// Wait for activation button
	_, err = ctx.page.WaitForSelector("button.activate", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "activation button not found")

	// Click activate button
	err = ctx.page.Click("button.activate")
	require.NoError(t, err, "failed to click activate button")

	// Wait for success message
	_, err = ctx.page.WaitForSelector(".success", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	require.NoError(t, err, "activation failed or timed out")
}

func (ctx *testContext) getLastError(name string) error {
	return ctx.lastErrors[name]
}

func (ctx *testContext) setLastError(name string, err error) {
	ctx.lastErrors[name] = err
}

func (ctx *testContext) clearAll() {
	// Navigate to the clear admin page
	_, err := ctx.page.Goto(ctx.frontendURL + "/admin/clear")
	if err != nil {
		return
	}

	// Wait for the clear button to be available
	_, err = ctx.page.WaitForSelector("button", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return
	}

	// Click the clear button
	err = ctx.page.Click("button")
	if err != nil {
		return
	}

	// Wait for success message to confirm clear completed
	_, err = ctx.page.WaitForSelector(".success", playwright.PageWaitForSelectorOptions{
		Timeout: playwright.Float(5000),
	})

	ctx.lastErrors = make(map[string]error)
}
