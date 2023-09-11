package models

import "database/sql"
type InvoiceInput struct {
	UserID	int     `json:"user_id"`
	Amount	float64 `json:"amount"`
	Label 	string  `json:"label"`
}

type InvoiceOutput struct {
	ID		int
	UserID 	int
	Amount 	float64
	Label  	string
	Status	sql.NullString
}
