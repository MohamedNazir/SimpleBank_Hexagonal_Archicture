package domain

// Customer Exported
type Customer struct {
	ID          int
	Name        string
	City        string
	Zipcode     int
	DateOfBirth string
	Status      string
}

//CustomerRepository Exported
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
