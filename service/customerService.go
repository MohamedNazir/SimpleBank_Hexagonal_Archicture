package service

import (
	"github.com/MohamedNazir/SimpleBank/domain"
	"github.com/MohamedNazir/SimpleBank/dto"
	errs "github.com/MohamedNazir/SimpleBank/errors"
)

//CustomerService exported
type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(ID int) (*dto.CustomerResponse, *errs.AppError)
}

//DefaultCustomerService exported
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

//GetAllCustomer exported
func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err
}

//GetCustomer exported
func (s DefaultCustomerService) GetCustomer(ID int) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

//NewCustomerService exported
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
