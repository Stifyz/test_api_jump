package data

import (
	"test_api/internal/models"
)

func UsersOutput() []models.UserOutput {
	return []models.UserOutput{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Balance:   1000,
		},
		{
			ID:        2,
			FirstName: "Alice",
			LastName:  "Smith",
			Balance:   2000,
		},
	}
}
