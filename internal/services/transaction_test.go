package services

import (
	"context"

	"testing"

	"github.com/stretchr/testify/assert"

	"test_api/internal/mocks"
	"test_api/internal/mocks/data"
)

func TestTransactionService_NewTransaction(t *testing.T) {
	// Create repository mock
	mockRepository := mocks.NewTransactionRepositoryMock()

	// Create service with repository's mock
	service := NewTransactionService(mockRepository)

	transaction := data.TransactionInput()

	// Define repository's mock behaviors
	mockRepository.On("NewTransaction", context.Background(), transaction).Return(nil)

	// Call service's method NewTransaction
	err := service.NewTransaction(context.Background(), transaction)

	// Assert results
	assert.NoError(t, err)

	mockRepository.AssertExpectations(t)
}
