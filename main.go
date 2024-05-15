package main

import (
	"database/sql"
	"log"

	"github.com/avijeetpandey/expense_tracker/api"
	db "github.com/avijeetpandey/expense_tracker/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:root@localhost:5432/expense_tracker?sslmode=disable"
	serverAddress = "0.0.0.0:9000"
)

func main() {
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(connection)

	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("unable to start the server", err)
	}
}
