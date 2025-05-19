package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// ConnectToDB connect to database
func ConnectToDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=tickets sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to DB")
	return db, nil
}
