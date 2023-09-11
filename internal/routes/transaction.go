package routes

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"test_api/internal/handlers"
	"test_api/internal/repositories"
	"test_api/internal/services"
)

func SetupoTransactionRoutes(r *mux.Router, db *sql.DB) {
	repository := repositories.NewTransactionRepository(db)
	service := services.NewTransactionService(repository)
	handler := handlers.NewTransactionHandler(service)

	r.HandleFunc(v1 + "/transaction", handler.NewTransaction()).Methods("POST")
}
