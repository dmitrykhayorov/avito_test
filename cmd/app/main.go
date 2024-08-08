package main

import (
	"avito/internal/api"
	"avito/internal/auth"
	"avito/internal/flat"
	"avito/internal/house"
)

func main() {
	authHandler := auth.NewAuthHandler()
	flatHandler := flat.NewFlatHandler()
	houseHandler := house.NewHouseHandler()

	server := api.NewServer(authHandler, flatHandler, houseHandler)
	server.Run(":8080")
}
