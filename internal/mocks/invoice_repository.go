package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"test_api/internal/models"
)

type InvoiceRepositoryMock struct {
	mock.Mock
}

func NewInvoiceRepositoryMock() *InvoiceRepositoryMock {
	return &InvoiceRepositoryMock{}
}

func (m *InvoiceRepositoryMock) CreateInvoice(ctx context.Context, invoice models.InvoiceInput) error {
	args := m.Called(ctx, invoice)
	return args.Error(0)
}
