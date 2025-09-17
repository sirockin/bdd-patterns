package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirockin/cucumber-screenplay-go/domain"
)

type HTTPClient struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string) *HTTPClient {
	return &HTTPClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (h *HTTPClient) CreateAccount(name string) error {
	reqBody := map[string]string{"name": name}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	resp, err := h.client.Post(h.baseURL+"/accounts", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("create account failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (h *HTTPClient) ClearAll() {
	req, err := http.NewRequest("DELETE", h.baseURL+"/clear", nil)
	if err != nil {
		return
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func (h *HTTPClient) GetAccount(name string) (domain.Account, error) {
	resp, err := h.client.Get(h.baseURL + "/accounts/" + name)
	if err != nil {
		return domain.Account{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return domain.Account{}, fmt.Errorf("Account not found: %s", name)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return domain.Account{}, fmt.Errorf("get account failed with status %d: %s", resp.StatusCode, string(body))
	}

	var account struct {
		Name          string `json:"name"`
		Activated     bool   `json:"activated"`
		Authenticated bool   `json:"authenticated"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Account{}, err
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		return domain.Account{}, err
	}

	// Create a domain account and set its fields using the accessor methods
	domainAccount := domain.NewAccount(account.Name)
	domainAccount.SetActivated(account.Activated)
	domainAccount.SetAuthenticated(account.Authenticated)

	return *domainAccount, nil
}

func (h *HTTPClient) Authenticate(name string) error {
	req, err := http.NewRequest("POST", h.baseURL+"/accounts/"+name+"/authenticate", nil)
	if err != nil {
		return err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("Account not found: %s", name)
	}

	if resp.StatusCode == http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		var errorResp struct {
			Error string `json:"error"`
		}
		json.Unmarshal(body, &errorResp)
		return fmt.Errorf(errorResp.Error)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("authenticate failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (h *HTTPClient) IsAuthenticated(name string) bool {
	resp, err := h.client.Get(h.baseURL + "/accounts/" + name + "/authentication-status")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	var authStatus struct {
		Authenticated bool `json:"authenticated"`
	}

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &authStatus)

	return authStatus.Authenticated
}

func (h *HTTPClient) Activate(name string) error {
	req, err := http.NewRequest("POST", h.baseURL+"/accounts/"+name+"/activate", nil)
	if err != nil {
		return err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("Account not found: %s", name)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("activate failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (h *HTTPClient) CreateProject(name string) error {
	req, err := http.NewRequest("POST", h.baseURL+"/accounts/"+name+"/projects", nil)
	if err != nil {
		return err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("Account not found: %s", name)
	}

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("create project failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (h *HTTPClient) GetProjects(name string) ([]domain.Project, error) {
	resp, err := h.client.Get(h.baseURL + "/accounts/" + name + "/projects")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Account not found: %s", name)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("get projects failed with status %d: %s", resp.StatusCode, string(body))
	}

	var projects []domain.Project
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

