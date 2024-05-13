package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDatabaseConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=expense_tracker sslmode=disable password=root host=localhost")

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected to the database")
	}

	return db, err
}

func GetDBInstance() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=expense_tracker sslmode=disable password=root host=localhost")
	if err != nil {
		return nil, err
	}
	return db, nil
}
