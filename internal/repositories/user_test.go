package repositories

import (
	"context"

	"testing"

	"github.com/stretchr/testify/assert"

	"test_api/internal/mocks"
	"test_api/internal/mocks/data"
)

func TestUserRepository_GetAllUsers(t *testing.T) {
	// Create DB mock
	mockDB, mock := mocks.NewMockDB(t)
	defer mockDB.Close()

	// Create repository with mocked DB
	repo := NewUserRepository(mockDB)

	// Define DB mock behaviour
	users := data.UsersOutput()
	mock.ExpectQuery("SELECT \\* FROM users").
		WillReturnRows(
			mock.NewRows([]string{"id", "first_name", "last_name", "balance"}).
				AddRow(users[0].ID, users[0].FirstName, users[0].LastName, users[0].Balance).
				AddRow(users[1].ID, users[1].FirstName, users[1].LastName, users[1].Balance),
		)

	// Call repository's GetAllUsers
	usersRes, err := repo.GetAllUsers(context.Background())

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, users, usersRes)

	assert.NoError(t, mock.ExpectationsWereMet())
}
