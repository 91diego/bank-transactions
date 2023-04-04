package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/91diego/bank-transactions/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

func (repo *PostgresRepository) InsertTransactions(ctx context.Context, trx *models.Transaction) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO transactions (id, transaction_date, transaction_amount, transaction_type, balance_id) VALUES ($1, $2, $3, $4, $5)", trx.ID, trx.TransactionDate, trx.TransactionAmount, trx.TransactionType, trx.BlanceID)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (repo *PostgresRepository) InsertBalance(ctx context.Context, balance *models.Balance) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO balances VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (id) DO UPDATE SET total = EXCLUDED.total + balances.total, debit_avarage = EXCLUDED.debit_avarage + balances.debit_avarage, credit_avarage = EXCLUDED.credit_avarage + balances.credit_avarage;",
		balance.ID, balance.Total, balance.DebitAvarage, balance.CreditAvarage, balance.Transactions, balance.UserID)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (repo *PostgresRepository) GetBalanceByID(ctx context.Context, balanceID string) (*models.Balance, error) {

	rows, err := repo.db.QueryContext(ctx, "SELECT total, debit_avarage, credit_avarage FROM balances WHERE id = $1", balanceID)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var balance = models.Balance{}
	for rows.Next() {
		if err = rows.Scan(&balance.Total, &balance.DebitAvarage, &balance.CreditAvarage); err == nil {
			return &balance, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &balance, nil
}
