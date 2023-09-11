package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"test_api/internal/models"
)

type UserServiceMock struct {
	mock.Mock
}

func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{}
}

func (m *UserServiceMock) GetAllUsers(ctx context.Context) ([]models.UserOutput, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.UserOutput), args.Error(1)
}
