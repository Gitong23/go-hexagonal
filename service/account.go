package service

import "bank/repository"

type NewAccountRequest struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type AccountResponse struct {
	AccountID   int     `json:"account_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type AccountService interface {
	NewAccount(int, NewAccountRequest) (*AccountResponse, error)
	GetAccounts(int) ([]AccountResponse, error)
}

func accRepoToRes(account *repository.Account) AccountResponse {
	return AccountResponse{
		AccountID:   account.AccountID,
		OpeningDate: account.OpeningDate,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.Status,
	}
}
