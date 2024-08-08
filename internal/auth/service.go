package auth

import (
	"avito/internal/tools"
	"avito/models"
	"context"
	"errors"
	"time"
)

const (
	expirePeriod time.Duration = time.Hour * time.Duration(3)
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) DummyLogin(ctx context.Context, userRole models.UserRole) (string, error) {
	switch userRole {
	case models.Client:
		break
	case models.Moderator:
		break
	default:
		return "", errors.New("wrong user_type: " + string(userRole))
	}
	token, err := tools.GenerateToken(string(userRole), expirePeriod)

	if err != nil {
		return "", err
	}

	return token, nil
}
