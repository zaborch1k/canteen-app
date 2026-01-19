package common

import (
	"errors"
	"log"
	"net/http"

	"canteen-app/internal/usecase"
)

func ErrorToHTTP(err error) (code int, msg string) {
	log.Println(err.Error())

	switch {
	case errors.Is(err, ErrInvalidRequest):
		return http.StatusBadRequest, "invalid request"

	case errors.Is(err, ErrValidationError):
		return http.StatusBadRequest, "validation error"

	case errors.Is(err, usecase.ErrInvalidCredentials):
		return http.StatusUnauthorized, "invalid credentials"

	case errors.Is(err, usecase.ErrInvalidRefresh),
		errors.Is(err, ErrRefreshTokenError):
		return http.StatusUnauthorized, "refresh token error"

	case errors.Is(err, usecase.ErrUserNotFound):
		return http.StatusNotFound, "user not found"

	case errors.Is(err, usecase.ErrUserExists):
		return http.StatusConflict, "user already exists"

	case errors.Is(err, usecase.ErrLoginInUse):
		return http.StatusConflict, "login already in use"

	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
