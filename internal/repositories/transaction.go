package repositories

import (
	"context"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"

	"test_api/internal/models"
	"test_api/internal/repositories/errors"
)

type ITransactionRepository interface {
	NewTransaction(context.Context, models.TransactionInput) error
}

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) NewTransaction(ctx context.Context, transaction models.TransactionInput) error {
	// Start transaction
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			// If error is returned cancel transaction
			tx.Rollback()
		}
	}()

	var invoiceID int
	var userID int
	var status sql.NullString
	var amount float64

	// Find Invoice
	err = tx.QueryRowContext(ctx, "SELECT id, user_id, amount, status FROM invoices WHERE id = $1", transaction.InvoiceID).
		Scan(&invoiceID, &userID, &amount, &status)

	if err != nil {
		return fmt.Errorf("errors: %w %w", errors.ErrTransaction_invoiceNotFound, err)
	}

	if amount != transaction.Amount {
		return errors.ErrTransaction_badAmount
	}

	badStatus := sql.NullString{String: "paid", Valid: true}
	if status == badStatus {
		return errors.ErrTransaction_badStatus
	}

	// Update user balance
	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", transaction.Amount, userID)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	// Update invoice status
	_, err = tx.ExecContext(ctx, "UPDATE invoices SET status = 'paid' WHERE id = $1", invoiceID)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	return nil
}
