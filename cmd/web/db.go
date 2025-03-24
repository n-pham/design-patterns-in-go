package main

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb"
)

// const (
// 	maxOpenDbConn = 25
// 	maxIdleDBConn = 25
// 	maxDBLifetime = 5 * time.Minute
// )

func initDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("duckdb", dsn)
	if err != nil {
		return nil, err
	}

	// test our database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// db.SetMaxOpenConns(maxOpenDbConn)
	// db.SetMaxIdleConns(maxIdleDBConn)
	// db.SetConnMaxLifetime(maxDBLifetime)

	return db, nil
}
