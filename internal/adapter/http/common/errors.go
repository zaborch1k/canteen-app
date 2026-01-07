package common

import "errors"

var (
	ErrInvalidRequest    = errors.New("invalid request")
	ErrRefreshTokenError = errors.New("no refresh token")
)
