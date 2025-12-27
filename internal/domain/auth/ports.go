package auth

import (
	"time"

	domUser "canteen-app/internal/domain/user"
)

type Claims struct {
	UserID    domUser.UserID
	Role      string
	ExpiresAt time.Time
}

type TokenService interface {
	GenerateAccesToken(c Claims) (string, error)
	ParseAccesToken(tokenStr string) (Claims, error)
}
