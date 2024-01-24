package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		// panic(err)
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createWordsTable := `
	CREATE TABLE IF NOT EXISTS words (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		mot TEXT NOT NULL,
		livre TEXT NOT NULL,
		page INTEGER NOT NULL,
		dateTime DATETIME NOT NULL
	)
	`
	_, err := DB.Exec(createWordsTable)
	if err != nil {
		panic("Could not create words table.")
	}
}
