package repository

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func Connect(dbName string) *sql.DB {
	// Connect to database
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		panic(err)
	}
	return db
}

func GetRows(db *sql.DB, query string) *sql.Rows {
	// Get rows from database
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	return rows
}

func GetRow(db *sql.DB, query string) *sql.Row {
	// Get row from database
	row := db.QueryRow(query)
	return row
}

func Exec(db *sql.DB, query string) sql.Result {
	// Execute query
	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	return res
}

func Prepare(db *sql.DB, query string) *sql.Stmt {
	// Prepare query
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	return stmt
}

func Close(db *sql.DB) {
	// Close database
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func CloseRows(rows *sql.Rows) {
	// Close rows
	err := rows.Close()
	if err != nil {
		panic(err)
	}
}
