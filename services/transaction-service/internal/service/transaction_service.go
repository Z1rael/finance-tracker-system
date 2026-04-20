package service

import (
	"finance-tracker-system/transaction-service/internal/model"
	"finance-tracker-system/transaction-service/internal/repository"
	"time"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

func (service *TransactionService) CreateTransaction(
	accountID int64,
	amount int64,
	description string,
	category int32,
	tType int32,
	timestamp time.Time,
) (*model.Transaction, error) {
	t := &model.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Description: description,
		Category:    category,
		Type:        tType,
		Timestamp:   timestamp,
	}

	return service.repo.Create(t)
}

func (service *TransactionService) ListTransactions(accountID int64) ([]*model.Transaction, error) {
	return service.repo.List(accountID)
}

func (service *TransactionService) DeleteTransaction(id int64) error {
	return service.DeleteTransaction(id)
}
