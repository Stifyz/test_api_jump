package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"test_api/internal/models"
)

type TransactionServiceMock struct {
	mock.Mock
}

func NewTransactionServiceMock() *TransactionServiceMock {
	return &TransactionServiceMock{}
}

func (m *TransactionServiceMock) NewTransaction(ctx context.Context, transaction models.TransactionInput) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}
