package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db}
}

func (a accountRepositoryDB) Create(acc Account) (*Account, error) {
	query := "insert into accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := a.db.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)
	return &acc, nil
}

func (a accountRepositoryDB) GetAll(customerId int) ([]Account, error) {

	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id = ?"
	accounts := []Account{}

	err := a.db.Select(&accounts, query, customerId)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
