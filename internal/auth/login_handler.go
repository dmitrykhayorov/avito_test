package auth

import (
	"avito/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: NewAuthService(),
	}
}

func (h *AuthHandler) DummyLogin(c *gin.Context) {
	userRole, ok := c.GetQuery("user_type")
	// TODO: setup proper logging
	if !ok {
		// TODO: add time to retry
		response := models.Response500{
			Message: "user type is not specified",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := h.service.DummyLogin(context.Background(), models.UserRole(userRole))
	if err != nil {
		response := models.Response500{
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.AutResponse200{
		Token: token,
	}

	c.JSON(http.StatusOK, response)
}
