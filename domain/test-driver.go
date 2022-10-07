package domain

import "fmt"

func NewDriver()*TestDriver{
	ret :=  TestDriver{}
	ret.ClearAll()
	return &ret
}

type TestDriver struct{
	accounts map[string]Account
	projects map[Account][]Project
}

func (d *TestDriver) ClearAll(){
	d.accounts=make(map[string]Account)
	d.projects=make(map[Account][]Project)
}

func( d *TestDriver) CreateAccount(name string)error{
	d.accounts[name]=Account{}
	return nil
}

func(d *TestDriver) GetAccount(name string)(Account,error){
	ret, present :=  d.accounts[name]
	if !present {
		return Account{}, fmt.Errorf("Account not found: %s", name)
	}
	return ret, nil
}

func (d *TestDriver) Authenticate(name string)error{
	return nil
}

func( d *TestDriver) IsAuthenticated(name string)bool{
	return false
}

func( d *TestDriver) GetProjects(name string)([]Project,error){
	account, err := d.GetAccount(name)
	if err != nil {
		return nil, err;
	}
	return d.projects[account], nil;
}

func( d *TestDriver) CreateProject(name string)error{
	account, err := d.GetAccount(name)
	if err != nil {
		return err;
	}
	d.projects[account]=append(d.projects[account], Project{})
	return nil
}