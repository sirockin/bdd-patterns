package main

func main(){}

func NewDomainDriver()*DomainDriver{
	return &DomainDriver{
		accounts: make(map[string]interface{}),
	}
}


type DomainDriver struct{
	accounts map[string]interface{}
}

func (d *DomainDriver) ClearAccounts(){
	d.accounts=make(map[string]interface{})
}

func( d *DomainDriver) CreateAccount(name string)error{
	d.accounts[name]=name
	return nil
}