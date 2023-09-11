package mocks

import (
    "database/sql"
    _ "github.com/lib/pq"
    
    "testing"
    "github.com/DATA-DOG/go-sqlmock"
)

func NewMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatal(err)
    }

    return db, mock
}
