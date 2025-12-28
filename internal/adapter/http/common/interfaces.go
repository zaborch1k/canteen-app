package common

import (
	domAuth "canteen-app/internal/domain/auth"
	domUser "canteen-app/internal/domain/user"
)

type AuthUseCase interface {
	Register(login, password, name, surname, role string) (*domAuth.Tokens, error)
	Login(login, password string) (*domAuth.Tokens, error)
	GetUserByLogin(login string) (*domUser.User, error)
	Refresh(refreshToken string) (*domAuth.Tokens, error)
	RevokeRefreshToken(refreshToken string) error
}
