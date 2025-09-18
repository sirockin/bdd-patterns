package domain

import (
	"fmt"

	"github.com/sirockin/cucumber-screenplay-go/internal/domain"
)

func New()*Domain{
	ret :=  Domain{}
	ret.ClearAll()
	return &ret
}

type Domain struct{
	accounts map[string]*domain.Account
	projects map[domain.Account][]domain.Project
}

func (d *Domain) ClearAll(){
	d.accounts=make(map[string]*domain.Account)
	d.projects=make(map[domain.Account][]domain.Project)
}

func( d *Domain) CreateAccount(name string)error{
	d.accounts[name]=domain.NewAccount(name)
	return nil
}

func(d *Domain) GetAccount(name string)(domain.Account,error){
	ret, present :=  d.accounts[name]
	if !present {
		return domain.Account{}, fmt.Errorf("Account not found: %s", name)
	}
	return *ret, nil
}

func (d *Domain) Activate(name string)error{
	account := d.accounts[name]
	if account == nil {
		return fmt.Errorf("Account not found: %s", name)
	}
	account.SetActivated(true)
	account.SetAuthenticated(true)  // Activation also authenticates the user
	return nil
}

func( d *Domain) IsActivated(name string)bool{
	account, err := d.GetAccount(name)
	if err != nil {
		return false
	}
	return account.IsActivated()
}

func (d *Domain) Authenticate(name string)error{
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

func( d *Domain) IsAuthenticated(name string)bool{
	account, err := d.GetAccount(name)
	if err != nil {
		return false
	}
	return account.IsAuthenticated()
}

func( d *Domain) GetProjects(name string)([]domain.Project,error){
	account, err := d.GetAccount(name)
	if err != nil {
		return nil, err;
	}
	return d.projects[account], nil;
}

func( d *Domain) CreateProject(name string)error{
	account, err := d.GetAccount(name)
	if err != nil {
		return err;
	}
	d.projects[account]=append(d.projects[account], domain.Project{})
	return nil
}