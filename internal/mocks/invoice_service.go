package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"test_api/internal/models"
)

type InvoiceServiceMock struct {
	mock.Mock
}

func NewInvoiceServiceMock() *InvoiceServiceMock {
	return &InvoiceServiceMock{}
}

func (m *InvoiceServiceMock) CreateInvoice(ctx context.Context, invoice models.InvoiceInput) error {
	args := m.Called(ctx, invoice)
	return args.Error(0)
}
