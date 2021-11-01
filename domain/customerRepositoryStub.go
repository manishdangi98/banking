package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Jon", City: "Ney York", DateofBirth: "2001-12-01", Status: "1"},
		{Id: "1002", Name: "Wiz", City: "San Jose", DateofBirth: "2001-11-01", Status: "1"},
		{Id: "1003", Name: "Alex", City: "San Francisco", DateofBirth: "2001-10-01", Status: "0"},
	}  
	return CustomerRepositoryStub{customers: customers}
}
