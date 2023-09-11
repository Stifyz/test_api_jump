package services

import (
	"context"

	"testing"

	"github.com/stretchr/testify/assert"

	"test_api/internal/mocks"
	"test_api/internal/mocks/data"
)

func TestUserService_GetAllUsers(t *testing.T) {
	// Create repository mock
	mockRepository := mocks.NewUserRepositoryMock()

	// Create service with repository's mock
	service := NewUserService(mockRepository)

	// Define repository's mock behaviors
	expectedUsers := data.UsersOutput()
	mockRepository.On("GetAllUsers", context.Background()).Return(expectedUsers, nil)

	// Call service's method GetAllUsers
	users, err := service.GetAllUsers(context.Background())

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)

	mockRepository.AssertExpectations(t)
}
