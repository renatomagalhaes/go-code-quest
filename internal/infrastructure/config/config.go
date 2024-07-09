package config

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    return db
}
