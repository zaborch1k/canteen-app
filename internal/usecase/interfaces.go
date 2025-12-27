package usecase

import (
	"time"

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

type AuthUseCase interface {
	Register(login, password, name, surname, role string) (*Tokens, error)
	Login(login, password string) (*Tokens, error)
	GetUserByLogin(login string) (*domUser.User, error)
	Refresh(refreshToken string) (*Tokens, error)
}
