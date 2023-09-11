package db

import (
    "fmt"

    "database/sql"
    _ "github.com/lib/pq"
)

// InitDB init connection to postgres
func InitDB() (*sql.DB, error) {
    // Create connection chain
    connStr := "host=db user=jump password=password dbname=jump sslmode=disable"

    // Open connection to database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Check if database connection is ok
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Connected to PostgreSQL database")

    return db, nil
}