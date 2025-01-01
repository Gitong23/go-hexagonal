package service

import (
	"bank/repository"
	"database/sql"
	"errors"
	"log"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerId: customer.CustomerId,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, customerResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}

		log.Println(err)
		return nil, err
	}
	// return nil, nil

	custResponse := CustomerResponse{
		CustomerId: customer.CustomerId,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
