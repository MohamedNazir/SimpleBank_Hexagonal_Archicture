package service

import (
	"github.com/MohamedNazir/SimpleBank/domain"
	errs "github.com/MohamedNazir/SimpleBank/errors"
)

//CustomerService exported
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(ID int) (*domain.Customer, *errs.AppError)
}

//DefaultCustomerService exported
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

//GetAllCustomer exported
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

//GetCustomer exported
func (s DefaultCustomerService) GetCustomer(ID int) (*domain.Customer, *errs.AppError) {
	return s.repo.FindByID(ID)
}

//NewCustomerService exported
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
