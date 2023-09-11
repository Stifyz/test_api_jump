package routes

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var v1 = "/v1"

func SetupRoutes(r *mux.Router, db *sql.DB) {
	SetupInvoiceRoutes(r, db)
	SetupoTransactionRoutes(r, db)
	SetupUserRoutes(r, db)
}