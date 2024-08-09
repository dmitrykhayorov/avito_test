package main

import (
	"avito/internal/api"
	"avito/internal/auth"
	"avito/internal/flat"
	"avito/internal/house"
	"avito/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	dbHost      string
	dbPort      string
	dbUser      string
	dbPassword  string
	dbName      string
	servicePort string
)

func loadEnvVariables() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	servicePort = os.Getenv("SERVICE_PORT")
}

func run() {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("cannot connect db")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot connect db")
	}

	houseRepository := repository.NewHouseRepository(db)
	flatRepository := repository.NewFlatRepository(db)

	authHandler := auth.NewAuthHandler()
	flatHandler := flat.NewFlatHandler(flatRepository)
	houseHandler := house.NewHouseHandler(houseRepository)

	server := api.NewServer(authHandler, flatHandler, houseHandler)
	server.Run(":" + servicePort)
}

func main() {
	run()
}
