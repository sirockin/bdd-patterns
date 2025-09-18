package domain

import "fmt"

type Project struct {
}

// Domain provides business operations for the domain
type Domain struct {
	accounts map[string]*Account
	projects map[Account][]Project
}

// New creates a new domain
func New() *Domain {
	d := &Domain{}
	d.ClearAll()
	return d
}

// ClearAll removes all data
func (d *Domain) ClearAll() {
	d.accounts = make(map[string]*Account)
	d.projects = make(map[Account][]Project)
}

// CreateAccount creates a new account
func (d *Domain) CreateAccount(name string) error {
	d.accounts[name] = NewAccount(name)
	return nil
}

// GetAccount retrieves an account by name
func (d *Domain) GetAccount(name string) (Account, error) {
	account, exists := d.accounts[name]
	if !exists {
		return Account{}, fmt.Errorf("Account not found: %s", name)
	}
	return *account, nil
}

// Activate activates an account and also authenticates the user
func (d *Domain) Activate(name string) error {
	account := d.accounts[name]
	if account == nil {
		return fmt.Errorf("Account not found: %s", name)
	}
	account.SetActivated(true)
	account.SetAuthenticated(true) // Activation also authenticates the user
	return nil
}

// IsActivated checks if an account is activated
func (d *Domain) IsActivated(name string) bool {
	account, err := d.GetAccount(name)
	if err != nil {
		return false
	}
	return account.IsActivated()
}

// Authenticate authenticates an account (requires activation first)
func (d *Domain) Authenticate(name string) error {
	account := d.accounts[name]
	if account == nil {
		return fmt.Errorf("Account not found: %s", name)
	}
	if !account.IsActivated() {
		return fmt.Errorf("%s, you need to activate your account", name)
	}
	account.SetAuthenticated(true)
	return nil
}

// IsAuthenticated checks if an account is authenticated
func (d *Domain) IsAuthenticated(name string) bool {
	account, err := d.GetAccount(name)
	if err != nil {
		return false
	}
	return account.IsAuthenticated()
}

// GetProjects retrieves projects for an account
func (d *Domain) GetProjects(name string) ([]Project, error) {
	account, err := d.GetAccount(name)
	if err != nil {
		return nil, err
	}
	return d.projects[account], nil
}

// CreateProject creates a project for an account
func (d *Domain) CreateProject(name string) error {
	account, err := d.GetAccount(name)
	if err != nil {
		return err
	}
	d.projects[account] = append(d.projects[account], Project{})
	return nil
}
