package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"

	"test_api/internal/db"
	"test_api/internal/routes"
)

func main() {
	// Initialize the database connection
	db, err := db.InitDB()
	if err != nil {
		fmt.Println("Error initializing the database:", err)
		return
	}
	defer db.Close()

	r := mux.NewRouter()
	routes.SetupRoutes(r, db)

	http.Handle("/", r)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
