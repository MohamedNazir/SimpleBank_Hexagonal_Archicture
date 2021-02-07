package domain

import (
	"github.com/MohamedNazir/SimpleBank/dto"
	errs "github.com/MohamedNazir/SimpleBank/errors"
)

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

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}
