package handlers

import (
	"context"

	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"

	"test_api/internal/mocks"
	"test_api/internal/mocks/data"
	"test_api/internal/repositories/errors"
)

func TestTransactionHandler_NewTransaction(t *testing.T) {
	// Use transaction test data
	transaction := data.TransactionInput()

	// Create a request body as JSON
	requestBody, err := json.Marshal(transaction)
	if err != nil {
		t.Fatal(err)
	}

	// Create simulate HTTP request
	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Create mocked service
	mockService := mocks.NewTransactionServiceMock()

	// Create handler with mocked service
	handler := NewTransactionHandler(mockService)

	// Define service behavior
	mockService.On("NewTransaction", context.Background(), transaction).Return(nil)

	// Call handler
	handler.NewTransaction()(w, req)

	// Assert status code of HTTP response
	assert.Equal(t, http.StatusNoContent, w.Code)

	assert.NoError(t, err)

	mockService.AssertExpectations(t)
}
func TestTransactionHandler_NewTransaction_ErrBadAmount(t *testing.T) {
	// Use transaction test data
	transaction := data.TransactionInput()

	// Create a request body as JSON
	requestBody, err := json.Marshal(transaction)
	if err != nil {
		t.Fatal(err)
	}

	// Create simulate HTTP request
	req := httptest.NewRequest("POST", "/transaction", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Create mocked service
	mockService := mocks.NewTransactionServiceMock()

	// Create handler with mocked service
	handler := NewTransactionHandler(mockService)

	// Define service behavior
	mockService.On("NewTransaction", context.Background(), transaction).Return(errors.ErrTransaction_badAmount)

	// Call handler
	handler.NewTransaction()(w, req)

	// Assert status code of HTTP response
	assert.Equal(t, http.StatusBadRequest, w.Code)

	assert.NoError(t, err)

	mockService.AssertExpectations(t)
}
func TestTransactionHandler_NewTransaction_ErrBadStatus(t *testing.T) {
	// Use transaction test data
	transaction := data.TransactionInput()

	// Create a request body as JSON
	requestBody, err := json.Marshal(transaction)
	if err != nil {
		t.Fatal(err)
	}

	// Create simulate HTTP request
	req := httptest.NewRequest("POST", "/transaction", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Create mocked service
	mockService := mocks.NewTransactionServiceMock()

	// Create handler with mocked service
	handler := NewTransactionHandler(mockService)

	// Define service behavior
	mockService.On("NewTransaction", context.Background(), transaction).Return(errors.ErrTransaction_badStatus)

	// Call handler
	handler.NewTransaction()(w, req)

	// Assert status code of HTTP response
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	assert.NoError(t, err)

	mockService.AssertExpectations(t)
}

func TestTransactionHandler_NewTransaction_ErrInvoiceNotFound(t *testing.T) {
	// Use transaction test data
	transaction := data.TransactionInput()

	// Create a request body as JSON
	requestBody, err := json.Marshal(transaction)
	if err != nil {
		t.Fatal(err)
	}

	// Create simulate HTTP request
	req := httptest.NewRequest("POST", "/transaction", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Create mocked service
	mockService := mocks.NewTransactionServiceMock()

	// Create handler with mocked service
	handler := NewTransactionHandler(mockService)

	// Define service behavior
	mockService.On("NewTransaction", context.Background(), transaction).Return(errors.ErrTransaction_invoiceNotFound)

	// Call handler
	handler.NewTransaction()(w, req)

	// Assert status code of HTTP response
	assert.Equal(t, http.StatusNotFound, w.Code)

	assert.NoError(t, err)

	mockService.AssertExpectations(t)
}