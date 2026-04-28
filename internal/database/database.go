package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Connect(dsn string) *sql.DB {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatalf("database: failed to open: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("database: failed to ping: %v", err)
	}

	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		log.Fatalf("database: failed to enable WAL mode: %v", err)
	}

	log.Println("database: connected with WAL mode enabled")
	return db
}
