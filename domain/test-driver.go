package domain

import "fmt"

func NewDriver()*TestDriver{
	return &TestDriver{
		accounts: make(map[string]Account),
	}
}

type TestDriver struct{
	accounts map[string]Account
}

func (d *TestDriver) ClearAccounts(){
	d.accounts=make(map[string]Account)
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