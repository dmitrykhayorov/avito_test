package tools

import (
	"avito/internal/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JwtCustomClaims struct {
	UserRole string
	jwt.RegisteredClaims
}

func GenerateToken(role string, expirePeriod time.Duration) (string, error) {
	// TODO: write proper token generation
	expireTimeNumeric := jwt.NewNumericDate(time.Now().Add(expirePeriod))
	claims := JwtCustomClaims{
		UserRole:         role,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: expireTimeNumeric},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_STRING")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GetRoleFromToken(tokenString string) (models.UserRole, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
		}
		return []byte(os.Getenv("SECRET_STRING")), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userRole := claims["UserRole"].(string)
		return models.UserRole(userRole), nil
	} else {
		fmt.Println("ok: ", ok, " valid token: ", token.Valid)
		return "", errors.New("token is invalid")
	}
}
