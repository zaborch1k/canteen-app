package api

import (
	"log"
	"net/http"
	"time"

	_ "canteen-app/cmd/docs"
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

type AccessTokenResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type RegisterRequest struct {
	Login    string `json:"login" binding:"required" example:"the_real_slim_shady"`
	Password string `json:"password" binding:"required" example:"password1234"`
	Name     string `json:"name" binding:"required" example:"Slim"`
	Surname  string `json:"surname" binding:"required" example:"Shady"`
	Role     string `json:"role" binding:"required" example:"admin"`
}

// Register godoc
//
//	@Summary		Регистрация пользователя
//	@Description	Создает нового пользователя, устанавливает refresh токен в cookie и возвращает access токен в теле ответа.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		RegisterRequest				true	"Данные для регистрации"
//	@Success		201		{object}	AccessTokenResponse			"Пользователь успешно зарегистрирован"
//	@Failure		400		{object}	InvalidRequestErrorResponse	"Некорректный запрос"
//	@Failure		409		{object}	UserExistsErrorResponse		"Пользователь с таким логином уже существует"
//	@Failure		500		{object}	InternalServerErrorResponse	"Внутренняя ошибка сервера"
//	@Router			/auth/register [post]
func (ah *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, ErrInvalidRequest)
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

	c.JSON(http.StatusCreated, AccessTokenResponse{AccessToken: tokens.AccessToken})
}

type LoginRequest struct {
	Login    string `json:"login" binding:"required" example:"the_real_slim_shady"`
	Password string `json:"password" binding:"required" example:"password1234"`
}

// Login godoc
//
//	@Summary		Аутентификация пользователя
//	@Description	Аутентифицирует существующего пользователя, устанавливает refresh токен в cookie и возвращает access токен в теле ответа.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		LoginRequest					true	"Данные для входа"
//	@Success		200		{object}	AccessTokenResponse				"Пользователь успешно аутентифицирован"
//	@Failure		400		{object}	InvalidRequestErrorResponse		"Некорректный запрос"
//	@Failure		401		{object}	InvalidCredentialsErrorResponse	"Логин/пароль некорректен"
//	@Failure		500		{object}	InternalServerErrorResponse		"Внутренняя ошибка сервера"
//	@Router			/auth/login [post]
func (ah *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, ErrInvalidRequest)
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

	c.JSON(http.StatusOK, AccessTokenResponse{AccessToken: tokens.AccessToken})
}

// Refresh godoc
//
//	@Summary		Обновление access токена
//	@Description	Проверяет refresh токен, установленный в cookie, и возврашает в теле ответа новый access токен
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	AccessTokenResponse			"access токен успешно обновлен"
//	@Failure		401	{object}	RefreshTokenErrorResponse	"Refresh токен не установлен или некорректен"
//	@Failure		500	{object}	InternalServerErrorResponse	"Внутренняя ошибка сервера"
//	@Router			/auth/refresh [get]
func (ah *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		writeError(c, ErrRefreshTokenError)
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
	c.JSON(http.StatusOK, AccessTokenResponse{AccessToken: tokens.AccessToken})
}

// Logout godoc
//
//	@Summary		Выход из системы
//	@Description	Инвалидирует refresh токен в cookie
//	@Tags			auth
//	@Success		204	"Успешный выход, тело ответа отсутствует"
//	@Router			/auth/logout [post]
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
