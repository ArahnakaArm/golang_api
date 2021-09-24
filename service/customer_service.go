package service

import (
	"database/sql"
	"hello/errs"
	"hello/logs"
	"hello/repository"
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
		/* 		log.Println(err) */
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
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
			return nil, errs.NewNotFoundError("Customer not found")
		}

		/* 	log.Println(err) */
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		CustomerId: customer.CustomerId,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
