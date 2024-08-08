package main

import (
	"avito/internal/api"
	"avito/internal/auth"
	"avito/internal/flat"
)

func main() {
	authHandler := auth.NewAuthHandler()
	flatHandler := flat.NewFlatHandler()

	server := api.NewServer(authHandler, flatHandler)
	server.Run(":8080")
}
