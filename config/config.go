package config

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func InitDatabase(connString string) *sql.DB {
    db, err := sql.Open("postgres", connString)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to the database: %v", err))
    }

    // Test the connection
    if err := db.Ping(); err != nil {
        panic(fmt.Sprintf("Database connection error: %v", err))
    }

    fmt.Println("Successfully connected to the database!")
    return db
}
