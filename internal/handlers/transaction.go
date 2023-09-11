package handlers

import (
	stdErrors "errors"

	"encoding/json"

	"net/http"

	"test_api/internal/models"
	"test_api/internal/repositories/errors"
	"test_api/internal/services"
)

type ITransactionHandler interface {
	NewTransaction() http.HandlerFunc
}

type TransactionHandler struct {
	service services.ITransactionService
}

func NewTransactionHandler(service services.ITransactionService) *TransactionHandler {
	return &TransactionHandler{service}
}

func (h *TransactionHandler) NewTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.TransactionInput

		// Decode and read json
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&transaction); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := h.service.NewTransaction(r.Context(), transaction)
		if err != nil {
			errorStatus := http.StatusInternalServerError
			if stdErrors.Is(err, errors.ErrTransaction_badStatus) {
				errorStatus = http.StatusUnprocessableEntity
			}
			if stdErrors.Is(err, errors.ErrTransaction_badAmount) {
				errorStatus = http.StatusBadRequest
			}
			if stdErrors.Is(err, errors.ErrTransaction_invoiceNotFound) {
				errorStatus = http.StatusNotFound
			}
			http.Error(w, err.Error(), errorStatus)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
