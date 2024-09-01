package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func createDatabase(db *sql.DB) {

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT,
		title TEXT,
		comment TEXT,
		repeat TEXT
	);
	`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	createIndexSQL := `
	CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
	`
	_, err = db.Exec(createIndexSQL)
	if err != nil {
		log.Fatal("Error creating index:", err)
	}

	log.Println("Table and index created successfully")
}
