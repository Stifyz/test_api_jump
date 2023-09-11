package services

import (
	"context"

	"test_api/internal/models"
	"test_api/internal/repositories"
)

type IUserService interface {
	GetAllUsers(context.Context) ([]models.UserOutput, error)
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) *UserService {
	return &UserService{repository}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.UserOutput, error) {
	users, err := s.repository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
