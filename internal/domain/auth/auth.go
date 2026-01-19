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

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
