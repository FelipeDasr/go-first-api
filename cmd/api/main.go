package main

import (
	"go-databases/internal/db"
	"go-databases/internal/httpserver"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file.\n\n%v", err)
	}

	db.StartConnection()
	defer db.Connection.Close();

	httpserver.Start()
}