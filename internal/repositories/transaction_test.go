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

func TestTransactionRepository_NewTransaction(t *testing.T) {
	// Create DB mock
	mockDB, mock := mocks.NewMockDB(t)
	defer mockDB.Close()

	// Create repository with mocked DB
	repo := NewTransactionRepository(mockDB)

	// Configure sql mock
	mock.ExpectBegin()

	invoice := data.InvoiceOutput_withUserID(false, 1)
	transaction := data.TransactionInput_withInvoice(invoice)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_id, amount, status FROM invoices WHERE id = $1")).
		WithArgs(invoice.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "amount", "status"}).
			AddRow(invoice.ID, invoice.UserID, invoice.Amount, invoice.Status),
		)
	
	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET balance = balance + $1 WHERE id = $2")).
		WithArgs(transaction.Amount, invoice.UserID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	
	mock.ExpectExec(regexp.QuoteMeta("UPDATE invoices SET status = 'paid' WHERE id = $1")).
		WithArgs(transaction.InvoiceID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	
	mock.ExpectCommit()

	err := repo.NewTransaction(context.Background(), transaction)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
