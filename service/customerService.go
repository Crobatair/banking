package service

import (
	"github.com/crobatair/banking/domain"
	"github.com/crobatair/banking/errs"
	"net/url"
)

// CustomerService Primary Interface, Any service must implement all methods to gain
// The trait of CustomerService:
// ie:
// 		type SomeService struct {
// 			anyVariable domain.AnyRepository
// 		}
//
//      func (s SomeService) FindAll() {}
//
// 		This will bind:
//			SomeService, to be an impl of CustomerService, because it implements all methods
//
type CustomerService interface {
	FindAllCustomers(url.Values) ([]domain.Customer, *errs.AppError)
	FindCustomerById(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService This struct, will define a repository for a CustomerRepository
// it's mandatory when instancing this DefaultService, to provide a repository.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// FindAllCustomers This method, will return all customers
// This method, all DefaultCustomerService to be an impl of CustomerService
// This allows, that any impl of CustomerService  to be bind to a DefaultCustomerService
func (s DefaultCustomerService) FindAllCustomers(f url.Values) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(f)
}

func (s DefaultCustomerService) FindCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.FindById(id)
}

// NewCustomerService This function, will return a new instance of DefaultCustomerService
// it will take an argument of repository and will return a new instance of DefaultCustomerService
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
