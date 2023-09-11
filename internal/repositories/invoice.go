package repositories

import (
	"fmt"

	"context"

	"database/sql"

	_ "github.com/lib/pq"

	"test_api/internal/models"
)

type IInvoiceRepository interface {
	CreateInvoice(context.Context, models.InvoiceInput) error
}

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db}
}

func (r *InvoiceRepository) CreateInvoice(ctx context.Context, invoice models.InvoiceInput) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO invoices (user_id, amount, label) VALUES ($1, $2, $3)",
		invoice.UserID, invoice.Amount, invoice.Label,
	)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}
