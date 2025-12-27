package usecase

import (
	"time"

	domAuth "canteen-app/internal/domain/auth"
	domUser "canteen-app/internal/domain/user"
)

type AuthRepository interface {
	CreateUser(user domUser.User) domUser.UserID
	GetUserByID(id domUser.UserID) (*domUser.User, error)
	GetUserByLogin(login string) (*domUser.User, error)
}

type RefreshTokenRepository interface {
	Save(tokenID string, userID domUser.UserID, exp time.Time)
	Delete(tokenID string)
	IsValid(tokenID string, userID domUser.UserID) bool
}

type TokenService interface {
	GenerateAccessToken(userID domUser.UserID, role string) (string, error)
	ParseAccessToken(tokenStr string) (domAuth.Claims, error)
	GenerateRefreshToken(userID domUser.UserID) (string, string, time.Time, error)
	ParseRefreshToken(tokenStr string) (domUser.UserID, string, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}
