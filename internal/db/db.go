package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/stdlib"
)

func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("DB close error: %v", err)
	}
}
