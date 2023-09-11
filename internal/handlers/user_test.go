package handlers

import (
	"context"

	"encoding/json"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"

	"test_api/internal/mocks"
	"test_api/internal/mocks/data"
	"test_api/internal/models"
)

func TestUserHandler_GetAllUsers(t *testing.T) {
	// Create simulate HTTP request
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	// Create mocked service
	mockService := mocks.NewUserServiceMock()

	// Create handler with mocked service
	handler := NewUserHandler(mockService)

	// Define service behavior
	expectedUsers := data.UsersOutput()
	mockService.On("GetAllUsers", context.Background()).Return(expectedUsers, nil)

	// Call handler
	handler.GetAllUsers()(w, req)

	// Assert status code of HTTP response
	assert.Equal(t, http.StatusOK, w.Code)

	// Extract json body
	var users []models.UserOutput
	err := json.NewDecoder(w.Body).Decode(&users)
	assert.NoError(t, err)

	// Asser response body
	assert.Equal(t, expectedUsers, users)

	mockService.AssertExpectations(t)
}
