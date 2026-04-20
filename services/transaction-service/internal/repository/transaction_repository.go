package repository

import (
	"finance-tracker-system/transaction-service/internal/model"
	"sync"
)

// TODO: Need to implement this for a database

type TransactionRepository interface {
	Create(t *model.Transaction) (*model.Transaction, error)
	List(accountID int64) ([]*model.Transaction, error)
	Delete(id int64) error
}

type InMemoryTransactionRepository struct {
	mu           sync.Mutex
	transactions map[int64]*model.Transaction
	nextID       int64
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{
		transactions: make(map[int64]*model.Transaction),
		nextID:       1,
	}
}

// TODO: error handling when db implementatin is done
func (repo *InMemoryTransactionRepository) Create(t *model.Transaction) (*model.Transaction, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	t.ID = repo.nextID
	repo.nextID++
	repo.transactions[t.ID] = t

	return t, nil
}

// TODO: error handling still missing
func (repo *InMemoryTransactionRepository) List(accountID int64) ([]*model.Transaction, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var transactions []*model.Transaction
	for _, t := range repo.transactions {
		if t.AccountID == accountID {
			transactions = append(transactions, t)
		}
	}

	return transactions, nil
}

// TODO: this does not work, easier to just implement with real db
func (repo *InMemoryTransactionRepository) Delete(id int64) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.transactions, id)

	return nil
}

// TODO: extend repo to also update transactions
