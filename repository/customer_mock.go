package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerId: 1001, Name: "ASH", City: "NR", ZipCode: "9512", Status: 1},
		{CustomerId: 1002, Name: "ASHD", City: "TNR", ZipCode: "1523", Status: 1},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerId == id {
			return &customer, nil
		}
	}
	return nil, errors.New("Customer not found")
}
