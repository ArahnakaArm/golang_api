package service

import (
	"database/sql"
	"errors"
	"hello/repository"
	"log"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	custResponses := []CustomerResponse{}

	for _, customers := range customers {
		custResponse := CustomerResponse{
			CustomerId: customers.CustomerId,
			Name:       customers.Name,
			Status:     customers.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {

	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("Customers not found")
		}

		log.Println(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerId: customer.CustomerId,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
