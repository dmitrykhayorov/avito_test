package api

import (
	"avito/internal/auth"
	"avito/internal/flat"
	"avito/internal/house"
	"avito/internal/models"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	router       *gin.Engine
	authHandler  *auth.AuthHandler
	flatHandler  *flat.FlatHandler
	houseHandler *house.HouseHandler
}

func NewServer(authHandler *auth.AuthHandler, flatHandler *flat.FlatHandler, houseHandler *house.HouseHandler) *Server {
	router := gin.Default()

	server := &Server{router: router}
	server.authHandler = authHandler
	server.flatHandler = flatHandler
	server.houseHandler = houseHandler

	server.routes()
	return server
}

func (s *Server) routes() {
	authRoutes := s.router.Group("/")
	authRoutes.GET("/dummyLogin", s.authHandler.DummyLogin)

	// TODO: implement
	flatRoutes := s.router.Group("/flat")
	flatRoutes.Use(auth.AuthMiddleware())
	{
		flatRoutes.POST("/create", s.flatHandler.Create)
		// TODO: create handler for updating a flat
		flatRoutes.POST("/update", auth.RoleMiddleware(models.Moderator))
	}

	// TODO: implement house logic
	houseRoutes := s.router.Group("/house")
	houseRoutes.Use(auth.AuthMiddleware())
	{
		houseRoutes.POST("/create", auth.RoleMiddleware(models.Moderator), s.houseHandler.Create)
		houseRoutes.GET("/:id", s.houseHandler.Get)
	}

}

func (s *Server) Run(addr string) {
	if err := s.router.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
