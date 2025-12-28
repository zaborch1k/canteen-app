package api

import (
	"log"
	"net/http"
	"time"

	"canteen-app/internal/adapter/http/common"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth       common.AuthUseCase
	refreshTTL time.Duration
}

func NewAuthHandler(router *gin.Engine, auth common.AuthUseCase, tokens usecase.TokenService, refreshTTL time.Duration) {
	handler := &AuthHandler{auth: auth, refreshTTL: refreshTTL}

	{
		auth := router.Group("/api/auth")
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.GET("/refresh", handler.Refresh)
	}
}

type registerRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	tokens, err := ah.auth.Register(req.Login, req.Password, req.Name, req.Surname, req.Role)
	if err != nil {
		writeError(c, err)
		return
	}

	c.SetCookieData(&http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(ah.refreshTTL),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusCreated, gin.H{"access_token": tokens.AccessToken})
}

type loginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	tokens, err := ah.auth.Login(req.Login, req.Password)
	if err != nil {
		writeError(c, err)
		return
	}

	c.SetCookieData(&http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(ah.refreshTTL),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken})
}

func (ah *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	tokens, err := ah.auth.Refresh(refreshToken)
	if err != nil {
		writeError(c, err)
		return
	}

	c.SetCookieData(&http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(ah.refreshTTL),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	c.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken})
}

func (ah *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err == nil && refreshToken != "" {
		if err := ah.auth.RevokeRefreshToken(refreshToken); err != nil {
			log.Printf("failed to revoke refresh token: %v", err)
		}
	}

	c.SetCookieData(&http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Domain:   "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.Status(http.StatusNoContent)
}
