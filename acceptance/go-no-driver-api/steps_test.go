package features_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/back-end/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func personHasCreatedAnAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	reqBody := map[string]string{"name": name}
	jsonBody, err := json.Marshal(reqBody)
	require.NoError(t, err)

	resp, err := ctx.client.Post(ctx.baseURL+"/accounts", "application/json", bytes.NewBuffer(jsonBody))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusCreated, resp.StatusCode, "create account should return 201")
}

func personHasSignedUp(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	personHasCreatedAnAccount(t, ctx, name)
	getAccount(t, ctx, name)
	personActivatesTheirAccount(t, ctx, name)
}

func getAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	resp, err := ctx.client.Get(ctx.baseURL + "/accounts/" + name)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.NotEqual(t, http.StatusNotFound, resp.StatusCode, "account should exist")
	require.Equal(t, http.StatusOK, resp.StatusCode, "get account should return 200")
}

func personShouldBeAuthenticated(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	resp, err := ctx.client.Get(ctx.baseURL + "/accounts/" + name + "/authentication-status")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var authStatus struct {
		Authenticated bool `json:"authenticated"`
	}

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	err = json.Unmarshal(body, &authStatus)
	require.NoError(t, err)

	assert.True(t, authStatus.Authenticated, "person %s should be authenticated", name)
}

func personShouldNotBeAuthenticated(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	resp, err := ctx.client.Get(ctx.baseURL + "/accounts/" + name + "/authentication-status")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var authStatus struct {
		Authenticated bool `json:"authenticated"`
	}

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	err = json.Unmarshal(body, &authStatus)
	require.NoError(t, err)

	assert.False(t, authStatus.Authenticated, "person %s should not be authenticated", name)
}

func personShouldNotSeeAnyProjects(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	resp, err := ctx.client.Get(ctx.baseURL + "/accounts/" + name + "/projects")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.NotEqual(t, http.StatusNotFound, resp.StatusCode, "account should exist")
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var projects []entities.Project
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	err = json.Unmarshal(body, &projects)
	require.NoError(t, err)

	assert.Empty(t, projects, "person %s should not see any projects", name)
}

func personShouldSeeTheirProject(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	resp, err := ctx.client.Get(ctx.baseURL + "/accounts/" + name + "/projects")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.NotEqual(t, http.StatusNotFound, resp.StatusCode, "account should exist")
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var projects []entities.Project
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	err = json.Unmarshal(body, &projects)
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

	req, err := http.NewRequest("POST", ctx.baseURL+"/accounts/"+name+"/authenticate", nil)
	if err != nil {
		ctx.setLastError(name, err)
		return
	}

	resp, err := ctx.client.Do(req)
	if err != nil {
		ctx.setLastError(name, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		ctx.setLastError(name, fmt.Errorf("account not found: %s", name))
		return
	}

	if resp.StatusCode == http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		var errorResp struct {
			Error string `json:"error"`
		}
		_ = json.Unmarshal(body, &errorResp)
		ctx.setLastError(name, fmt.Errorf("%s", errorResp.Error))
		return
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		ctx.setLastError(name, fmt.Errorf("authenticate failed with status %d: %s", resp.StatusCode, string(body)))
		return
	}

	ctx.setLastError(name, nil)
}

func personCreatesAProject(t *testing.T, ctx *testContext, name string) {
	t.Helper()

	req, err := http.NewRequest("POST", ctx.baseURL+"/accounts/"+name+"/projects", nil)
	require.NoError(t, err)

	resp, err := ctx.client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.NotEqual(t, http.StatusNotFound, resp.StatusCode, "account should exist")
	require.Equal(t, http.StatusCreated, resp.StatusCode, "create project should return 201")
}

func personActivatesTheirAccount(t *testing.T, ctx *testContext, name string) {
	t.Helper()
	getAccount(t, ctx, name)

	req, err := http.NewRequest("POST", ctx.baseURL+"/accounts/"+name+"/activate", nil)
	require.NoError(t, err)

	resp, err := ctx.client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.NotEqual(t, http.StatusNotFound, resp.StatusCode, "account should exist")
	require.Equal(t, http.StatusOK, resp.StatusCode, "activate should return 200")
}

func (ctx *testContext) getLastError(name string) error {
	return ctx.lastErrors[name]
}

func (ctx *testContext) setLastError(name string, err error) {
	ctx.lastErrors[name] = err
}

func (ctx *testContext) clearAll() {
	req, err := http.NewRequest("DELETE", ctx.baseURL+"/clear", nil)
	if err != nil {
		return
	}

	resp, err := ctx.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	ctx.lastErrors = make(map[string]error)
}
