package main

import (
	"avito/internal/api"
	"avito/internal/auth"
	"avito/internal/flat"
	"avito/internal/house"
	"avito/internal/repository"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5443 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("cannot connect db")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot connect db")
	}
	houseRepository := repository.NewHouseRepository(db)

	authHandler := auth.NewAuthHandler()
	flatHandler := flat.NewFlatHandler()
	houseHandler := house.NewHouseHandler(houseRepository)

	server := api.NewServer(authHandler, flatHandler, houseHandler)
	server.Run(":8080")
}
