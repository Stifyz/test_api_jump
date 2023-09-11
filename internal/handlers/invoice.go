package handlers

import (
	"encoding/json"

	"net/http"

	"test_api/internal/models"
	"test_api/internal/services"
)

type IInvoiceHandler interface {
	CreateInvoice() http.HandlerFunc
}

type InvoiceHandler struct {
	service services.IInvoiceService
}

func NewInvoiceHandler(service services.IInvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

func (h *InvoiceHandler) CreateInvoice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var invoice models.InvoiceInput

		// Decode and read json
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&invoice); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := h.service.CreateInvoice(r.Context(), invoice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
