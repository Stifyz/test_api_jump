package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"test_api/internal/models"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) GetAllUsers(ctx context.Context) ([]models.UserOutput, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.UserOutput), args.Error(1)
}
