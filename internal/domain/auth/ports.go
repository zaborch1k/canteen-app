package auth

import (
	"time"

	"canteen-app/internal/domain"
)

type Claims struct {
	UserID   domain.UserID
	Role     string
	ExpireAt time.Time
}
