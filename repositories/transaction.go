package repositories

import (
	"context"

	"github.com/91diego/bank-transactions/models"
)

type TransactionRepository interface {
	Close() error
	InsertTransactions(ctx context.Context, user *models.Transaction) error
	InsertBalance(ctx context.Context, user *models.Balance) error
	GetBalanceByID(ctx context.Context, balanceID string) (*models.Balance, error)
}

var implementation TransactionRepository

func SetRepository(repository TransactionRepository) {
	implementation = repository
}

// Close db connection
func Close() error {
	return implementation.Close()
}

// InsertTransactions store new transactions in database
func InsertTransactions(ctx context.Context, trx *models.Transaction) error {
	return implementation.InsertTransactions(ctx, trx)
}

// InsertBalance store new balance in database
func InsertBalance(ctx context.Context, balance *models.Balance) error {
	return implementation.InsertBalance(ctx, balance)
}

// GetBalanceByID retrieve balanc by ID
func GetBalanceByID(ctx context.Context, balanceID string) (*models.Balance, error) {
	return implementation.GetBalanceByID(ctx, balanceID)
}
