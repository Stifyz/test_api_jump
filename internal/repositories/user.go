package repositories

import (
	"fmt"

	"context"

	"database/sql"

	"test_api/internal/models"
)

type IUserRepository interface {
	GetAllUsers(context.Context) ([]models.UserOutput, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.UserOutput, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var users []models.UserOutput
	for rows.Next() {
		var user models.UserOutput
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Balance)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over user rows: %w", err)
	}

	return users, nil
}
