package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerId: 1, Name: "John Doe", DateOfBirth: "1990-01-01", City: "Los Angeles", ZipCode: "90001", Status: "active"},
		{CustomerId: 2, Name: "Kevin", DateOfBirth: "1995-01-01", City: "Bangkok", ZipCode: "90001", Status: "inactive"},
	}
	return customerRepositoryMock{customers}
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

	return nil, errors.New("customer not found")
}
