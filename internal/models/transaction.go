package models

type TransactionInput struct {
	InvoiceID 	int		`json:"invoice_id"`
	Amount 		float64	`json:"amount"`
	Reference  	string	`json:"reference"`
}
