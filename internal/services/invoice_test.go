package services

import (
	"context"

	"testing"

	"github.com/stretchr/testify/assert"

	"test_api/internal/mocks"
	"test_api/internal/mocks/data"
)

func TestInvoiceService_CreateInvoice(t *testing.T) {
	// Create repository mock
	mockRepository := mocks.NewInvoiceRepositoryMock()

	// Create service with repository's mock
	service := NewInvoiceService(mockRepository)

	invoice := data.InvoiceInput()

	// Define repository's mock behaviors
	mockRepository.On("CreateInvoice", context.Background(), invoice).Return(nil)

	// Call service's method CreateInvoice
	err := service.CreateInvoice(context.Background(), invoice)

	// Assert results
	assert.NoError(t, err)

	mockRepository.AssertExpectations(t)
}
