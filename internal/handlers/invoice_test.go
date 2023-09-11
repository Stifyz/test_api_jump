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
)

func TestInvoiceHandler_CreateInvoice(t *testing.T) {
	// Use invoice test data
	invoice := data.InvoiceInput()

	// Create a request body as JSON
	requestBody, err := json.Marshal(invoice)
	if err != nil {
		t.Fatal(err)
	}

	// Create simulate HTTP request
	req := httptest.NewRequest("POST", "/invoice", bytes.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Create mocked service
	mockService := mocks.NewInvoiceServiceMock()

	// Create handler with mocked service
	handler := NewInvoiceHandler(mockService)

	// Define service behavior
	mockService.On("CreateInvoice", context.Background(), invoice).Return(nil)

	// Call handler
	handler.CreateInvoice()(w, req)

	// Assert status code of HTTP response
	assert.Equal(t, http.StatusNoContent, w.Code)

	assert.NoError(t, err)

	mockService.AssertExpectations(t)
}
