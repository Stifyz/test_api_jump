package routes

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"test_api/internal/handlers"
	"test_api/internal/repositories"
	"test_api/internal/services"
)

func SetupInvoiceRoutes(r *mux.Router, db *sql.DB) {
	repository := repositories.NewInvoiceRepository(db)
	service := services.NewInvoiceService(repository)
	handler := handlers.NewInvoiceHandler(service)

	r.HandleFunc(v1 + "/invoice", handler.CreateInvoice()).Methods("POST")
}
