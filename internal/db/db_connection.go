package db

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Connection is the connection with the database
var Connection *sql.DB;

// StartDBConnection starts a connection to the database
func StartConnection() {
	databaseURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		defer db.Close()
		log.Fatalf("Error opening database connection.\n\n%v", err)
	}

	ctx := context.Background()
	Connection = db

	if err = Connection.PingContext(ctx); err != nil {
		defer Connection.Close()
		log.Fatalf("Error pinging database.\n\n%v", err)
	}
}

// CreateQueryAndContext creates a new query and context
func CreateQueryAndContext() (*Queries, context.Context) {
	return New(Connection), context.Background()
}