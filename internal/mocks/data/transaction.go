package data

import (
	"test_api/internal/models"
)

func TransactionInput() models.TransactionInput {
	return models.TransactionInput{
		InvoiceID: 1,
		Amount:    100.0,
		Reference: "Transaction Reference",
	}
}

func TransactionInput_withInvoice(invoice models.InvoiceOutput) models.TransactionInput {
	return models.TransactionInput{
		InvoiceID: invoice.ID,
		Amount:    invoice.Amount,
		Reference: "Transaction Reference",
	}
}
