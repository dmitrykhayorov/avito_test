package api

import (
	"avito/internal/auth"
	"avito/internal/flat"
	"avito/internal/house"
	"avito/internal/models"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Server struct {
	Router       *gin.Engine
	authHandler  *auth.AuthHandler
	flatHandler  *flat.FlatHandler
	houseHandler *house.HouseHandler
}

func NewServer(authHandler *auth.AuthHandler, flatHandler *flat.FlatHandler, houseHandler *house.HouseHandler) *Server {
	router := gin.Default()

	server := &Server{Router: router}
	server.authHandler = authHandler
	server.flatHandler = flatHandler
	server.houseHandler = houseHandler

	server.routes()
	return server
}

func (s *Server) routes() {
	authRoutes := s.Router.Group("/")
	authRoutes.GET("/dummyLogin", s.authHandler.DummyLogin)

	flatRoutes := s.Router.Group("/flat")
	flatRoutes.Use(auth.AuthMiddleware())
	{
		flatRoutes.POST("/create", s.flatHandler.Create)
		flatRoutes.POST("/update", auth.RoleMiddleware(models.Moderator), s.flatHandler.Update)
	}

	houseRoutes := s.Router.Group("/house")
	houseRoutes.Use(auth.AuthMiddleware())
	{
		houseRoutes.POST("/create", auth.RoleMiddleware(models.Moderator), s.houseHandler.Create)
		houseRoutes.GET("/:id", s.houseHandler.GetFlatsByHouseID)
	}

}

func (s *Server) Run(addr string) {
	if err := s.Router.Run(addr); err != nil {
		slog.Error("failed to run server: " + err.Error())
		return
	}
}
