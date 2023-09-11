package routes

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"test_api/internal/handlers"
	"test_api/internal/repositories"
	"test_api/internal/services"
)

func SetupUserRoutes(r *mux.Router, db *sql.DB) {
	repository := repositories.NewUserRepository(db)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)

	r.HandleFunc(v1 + "/users", handler.GetAllUsers()).Methods("GET")
}
