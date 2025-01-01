package service

type CustomerResponse struct {
	CustomerId int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
