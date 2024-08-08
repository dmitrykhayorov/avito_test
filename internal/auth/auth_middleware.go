package auth

import (
	"avito/internal/tools"
	"avito/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			c.Abort()
			return
		}

		role, err := tools.GetRoleFromToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("userRole", string(role))
		fmt.Println("added userRole to context")
		c.Next()
	}
}

func RoleMiddleware(requiredRole models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exist := c.Get("userRole")

		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role is not found"})
			c.Abort()
			return
		}
		userRoleString := userRole.(string)
		if models.UserRole(userRoleString) != requiredRole {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access forbidden: insufficient permissions"})
			c.Abort()
			return
		}
		fmt.Println("role validated")
		c.Next()
	}
}
