package api

import (
	"canteen-app/internal/domain"
	"canteen-app/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	users usecase.UserUseCase
}

func NewAuthHandler(router *gin.Engine, users usecase.UserUseCase) {
	handler := &AuthHandler{users: users}

	{
		auth := router.Group("/api/auth")
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
}

type registerRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type registerResponse struct {
	AccessToken string `json:"access_token"`
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if _, err := ah.users.GetUserByLogin(req.Login); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user with this login has already been registered"})
		return
	}

	hash, err := domain.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	ah.users.RegisterUser(req.Login, hash, req.Name, req.Surname, req.Role)

	resp := registerResponse{AccessToken: "dfsf"}
	c.JSON(http.StatusCreated, resp)
}

type loginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := ah.users.GetUserByLogin(req.Login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := domain.CheckPassword(user.PasswordHash, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	resp := loginResponse{AccessToken: "dfsf"}
	c.JSON(http.StatusOK, resp)
}
