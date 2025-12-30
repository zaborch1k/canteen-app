package api

import (
	"errors"
	"log"
	"net/http"

	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func writeError(c *gin.Context, err error) {
	log.Println(err.Error())

	switch {
	case errors.Is(err, ErrInvalidRequest):
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})

	case errors.Is(err, usecase.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})

	case errors.Is(err, usecase.ErrInvalidRefresh),
		errors.Is(err, ErrRefreshTokenError):
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token error"})

	case errors.Is(err, usecase.ErrUserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})

	case errors.Is(err, usecase.ErrUserExists):
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
