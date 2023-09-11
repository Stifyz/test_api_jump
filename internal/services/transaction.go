package services

import (
	"context"

	"test_api/internal/models"
	"test_api/internal/repositories"
)

type ITransactionService interface {
	NewTransaction(context.Context, models.TransactionInput) error
}

type TransactionService struct {
	repository repositories.ITransactionRepository
}

func NewTransactionService(repository repositories.ITransactionRepository) *TransactionService {
	return &TransactionService{repository}
}

func (s *TransactionService) NewTransaction(ctx context.Context, transaction models.TransactionInput) error {
	err := s.repository.NewTransaction(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}
