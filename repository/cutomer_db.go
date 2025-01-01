package repository

import "github.com/jmoiron/sqlx"

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{db: db}
}

func (r CustomerRepositoryDB) GetAll() ([]Customer, error) {

	customers := []Customer{}
	query := `select customer_id, name, date_of_birth, city, zip_code, status from customers`

	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r CustomerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := `select customer_id, name, date_of_birth, city, zip_code, status from customers where customer_id = ?`
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
