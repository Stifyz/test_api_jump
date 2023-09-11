package repositories

import (
	"context"

	"regexp"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"test_api/internal/mocks"
	"test_api/internal/mocks/data"
)

func TestInvoiceRepository_CreateInvoice(t *testing.T) {
	// Create DB mock
	mockDB, mock := mocks.NewMockDB(t)
	defer mockDB.Close()

	// Create repository with mocked DB
	repo := NewInvoiceRepository(mockDB)

	// Define DB mock behaviour
	invoice := data.InvoiceInput()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO invoices (user_id, amount, label) VALUES ($1, $2, $3)")).WithArgs(
		invoice.UserID, invoice.Amount, invoice.Label,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateInvoice(context.Background(), invoice)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
