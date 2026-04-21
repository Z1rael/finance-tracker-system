package repository

import (
	"context"
	"finance-tracker-system/transaction-service/internal/model"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewPostgresPool(connString string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) Create(
	ctx context.Context,
	t *model.Transaction,
) (*model.Transaction, error) {
	transaction, err := repo.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback(ctx)

	// insert transaction
	query := `
		INSERT INTO transactions (account_id, amount, description, category, transaction_type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	err = transaction.QueryRow(ctx, query,
		t.AccountID,
		t.Amount,
		t.Description,
		t.Category,
		t.Type,
	).Scan(&t.ID, &t.Timestamp)

	if err != nil {
		return nil, err
	}

	// update account
	_, err = transaction.Exec(ctx, `
		UPDATE accounts
		SET balance = balance + $1,
		updated_at = NOW()
		WHERE id = $2
	`, t.Amount, t.AccountID)

	if err != nil {
		return nil, err
	}

	// commit transaction
	if err := transaction.Commit(ctx); err != nil {
		return nil, err
	}

	return t, nil
}

func (repo *TransactionRepository) Transfer(
	ctx context.Context,
	fromAccountId int64,
	toAccountId int64,
	amount int64,
	description string,
	category int32,
) error {
	transaction, err := repo.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer transaction.Rollback(ctx)

	// deduct from source
	_, err = transaction.Exec(ctx, `
		INSERT INTO transactions (account_id, amount, description, category, transaction_type)
		VALUES ($1, $2, $3, $4, $5)
	`, fromAccountId, -amount, description, category, 3) // Transfer

	if err != nil {
		return err
	}

	// add to destination
	_, err = transaction.Exec(ctx, `
		INSERT INTO transactions (account_id, amount, description, category, transaction_type)
		VALUES ($1, $2, $3, $4, $5)
	`, toAccountId, amount, description, category, 3)

	if err != nil {
		return err
	}

	// update accounts
	_, err = transaction.Exec(ctx, `
		UPDATE accounts SET balance = balance - $1,
		updated_at = NOW()
		WHERE id = $2
	`, amount, fromAccountId)
	if err != nil {
		return err
	}

	_, err = transaction.Exec(ctx, `
		UPDATE accounts SET balance = balance + $1,
		updated_at = NOW()
		WHERE id = $2
	`, amount, toAccountId)
	if err != nil {
		return err
	}

	return transaction.Commit(ctx)
}

func (repo *TransactionRepository) List(
	ctx context.Context,
	accountID int64,
) ([]*model.Transaction, error) {

	transaction, err := repo.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback(ctx)

	// get List
	rows, err := transaction.Query(ctx, `
		SELECT id, account_id, amount, description, category, transaction_type, created_at
		FROM transactions
		WHERE account_id = $1
		ORDER BY created_at DESC
	`, accountID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.Transaction

	for rows.Next() {
		var t model.Transaction
		err := rows.Scan(
			&t.ID,
			&t.AccountID,
			&t.Amount,
			&t.Description,
			&t.Category,
			&t.Type,
			&t.Timestamp,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, &t)
	}

	return result, nil
}

func (repo *TransactionRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	transaction, err := repo.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer transaction.Rollback(ctx)

	_, err = transaction.Exec(ctx, `
		DELETE FROM transactions WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	return transaction.Commit(ctx)
}

// TODO: extend repo to also update transactions
