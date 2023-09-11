package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"test_api/internal/models"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func NewTransactionRepositoryMock() *TransactionRepositoryMock {
	return &TransactionRepositoryMock{}
}

func (m *TransactionRepositoryMock) NewTransaction(ctx context.Context, transaction models.TransactionInput) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}
