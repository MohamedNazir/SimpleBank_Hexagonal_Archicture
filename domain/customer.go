package domain

import errs "github.com/MohamedNazir/SimpleBank/errors"

// Customer Exported
type Customer struct {
	ID          int `db:"customer_id"`
	Name        string
	City        string
	Zipcode     int
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

//CustomerRepository Exported
type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindByID(int) (*Customer, *errs.AppError)
}
