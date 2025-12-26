package auth

import (
	"time"

	domUser "canteen-app/internal/domain/user"
)

type Claims struct {
	UserID   domUser.UserID
	Role     string
	ExpireAt time.Time
}
