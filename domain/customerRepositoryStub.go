package domain

//CustomerRepositoryStub Exported
type CustomerRepositoryStub struct {
	customers []Customer
}

//FindAll Exported
func (stub CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return stub.customers, nil
}

//NewCustomerRepositoryStub exported
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customerslice := []Customer{
		{ID: 1001, Name: "Rizan", City: "Chennai", Zipcode: 600041, DateOfBirth: "2000-01-01", Status: "Active"},
		{ID: 1002, Name: "Mohamed", City: "New Delhi", Zipcode: 400230, DateOfBirth: "1990-02-04", Status: "Active"},
	}
	return CustomerRepositoryStub{customers: customerslice}
}
