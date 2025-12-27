package api

import (
	"net/http"

	"canteen-app/internal/adapter/http/common"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth common.AuthUseCase
}

func NewAuthHandler(router *gin.Engine, auth common.AuthUseCase, tokens usecase.TokenService) {
	handler := &AuthHandler{auth: auth}

	{
		auth := router.Group("/api/auth")
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/refresh", handler.Refresh)
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

	c.JSON(http.StatusCreated, tokens)
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

	c.JSON(http.StatusOK, tokens)
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (ah *AuthHandler) Refresh(c *gin.Context) {
	var req refreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	tokens, err := ah.auth.Refresh(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
	}
	c.JSON(http.StatusOK, tokens)
}
