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
	"time"

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
	loadEnvVariables()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	retries := 5
	var db *sql.DB
	var err error

	for i := 0; i < retries; i++ {
		db, err = sql.Open("postgres", connStr)

		err = db.Ping()
		if err == nil {
			break
		}
		fmt.Println("SLEEEEEEEEPPPPPPPINNNNGGGGG")
		time.Sleep(time.Second * 5)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("cannot connect db: ", err)
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
