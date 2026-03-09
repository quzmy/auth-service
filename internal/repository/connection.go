package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewDB(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
	return db
}
