package usecase

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
	ErrLoginInUse         = errors.New("login already in use")
	ErrUserNotFound       = errors.New("user not found")
	ErrMenuNotFound       = errors.New("menu not found")
	ErrInvalidRefresh     = errors.New("invalid refresh token")
)
