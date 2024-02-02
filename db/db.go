package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDb() {
	dbi, err := sql.Open("sqlite3", "api.db")
	db = dbi
	if err != nil {
		panic("Could not connect to the DB")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	_, err := db.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}
}
