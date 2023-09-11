package data

import (
	"database/sql"
	"test_api/internal/models"
)

func InvoiceInput() models.InvoiceInput {
	return models.InvoiceInput{
		UserID: 2,
		Amount: 120.11,
		Label:  "Test",
	}
}

func InvoiceOutput(paid bool) models.InvoiceOutput {
	var statusValue string
	if (paid) {
		statusValue = "paid"
	} else {
		statusValue = "pending"
	}
	return models.InvoiceOutput{
		ID: 1,
		UserID: 2,
		Amount: 120.11,
		Label:  "Test",
		Status: sql.NullString{String: statusValue, Valid: true},
	}
}

func InvoiceOutput_withUserID(paid bool, userID int) models.InvoiceOutput {
	var statusValue string
	if (paid) {
		statusValue = "paid"
	} else {
		statusValue = "pending"
	}
	return models.InvoiceOutput{
		ID: userID,
		UserID: 2,
		Amount: 120.11,
		Label:  "Test",
		Status: sql.NullString{String: statusValue, Valid: true},
	}
}
