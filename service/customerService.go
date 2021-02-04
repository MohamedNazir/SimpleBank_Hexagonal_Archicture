package service

import (
	"github.com/MohamedNazir/SimpleBank/domain"
)

//CustomerService exported
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

//DefaultCustomerService exported
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

//GetAllCustomer exported
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

//NewCustomerService exported
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
