package services

import (
	"context"

	"test_api/internal/models"
	"test_api/internal/repositories"
)

type IInvoiceService interface {
	CreateInvoice(context.Context, models.InvoiceInput) error
}

type InvoiceService struct {
	repository repositories.IInvoiceRepository
}

func NewInvoiceService(repository repositories.IInvoiceRepository) *InvoiceService {
	return &InvoiceService{repository}
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, invoice models.InvoiceInput) error {
	err := s.repository.CreateInvoice(ctx, invoice)
	if err != nil {
		return err
	}
	return nil
}
