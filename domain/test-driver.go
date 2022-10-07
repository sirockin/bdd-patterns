package domain

import "fmt"

func New()*Domain{
	ret :=  Domain{}
	ret.ClearAll()
	return &ret
}

type Domain struct{
	accounts map[string]*Account
	projects map[Account][]Project
}

func (d *Domain) ClearAll(){
	d.accounts=make(map[string]*Account)
	d.projects=make(map[Account][]Project)
}

func( d *Domain) CreateAccount(name string)error{
	d.accounts[name]=&Account{name:name}
	return nil
}

func(d *Domain) GetAccount(name string)(Account,error){
	ret, present :=  d.accounts[name]
	if !present {
		return Account{}, fmt.Errorf("Account not found: %s", name)
	}
	return *ret, nil
}

func (d *Domain) Activate(name string)error{
	account := d.accounts[name]
	if account == nil {
		return fmt.Errorf("Account not found: %s", name)
	}
	account.activated=true
	return nil
}

func( d *Domain) IsActivated(name string)bool{
	account, err := d.GetAccount(name)
	if err != nil {
		return false
	}
	return account.activated
}

func (d *Domain) Authenticate(name string)error{
	account := d.accounts[name]
	if account == nil {
		return fmt.Errorf("Account not found: %s", name)
	}
	if !account.activated {
		return fmt.Errorf("%s, you need to activate your account", name)
	}
	account.authenticated=true
	return nil
}

func( d *Domain) IsAuthenticated(name string)bool{
	account, err := d.GetAccount(name)
	if err != nil {
		return false
	}
	return account.activated
}

func( d *Domain) GetProjects(name string)([]Project,error){
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
	d.projects[account]=append(d.projects[account], Project{})
	return nil
}