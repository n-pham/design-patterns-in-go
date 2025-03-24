package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/marcboeker/go-duckdb"
)

const (
	maxOpenDbConn = 1
	maxIdleDBConn = 1
	maxDBLifetime = 5 * time.Minute
)

func initDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("duckdb", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// test our database
	if err = db.Ping(); err != nil {
		db.Close() // Ensure the database connection is closed if Ping fails
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)

	return db, nil
}
