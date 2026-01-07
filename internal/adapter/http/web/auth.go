package web

import (
	"errors"
	"net/http"
	"time"

	"canteen-app/internal/adapter/http/common"
	domUser "canteen-app/internal/domain/user"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth       common.AuthUseCase
	accessTTL  time.Duration
	refreshTTL time.Duration
	tokenSvc   usecase.TokenService
}

func NewAuthHandler(router *gin.Engine, auth common.AuthUseCase, accessTTL time.Duration, refreshTTL time.Duration, tokenSvc usecase.TokenService) {
	handler := &AuthHandler{auth: auth, accessTTL: accessTTL, refreshTTL: refreshTTL, tokenSvc: tokenSvc}

	router.LoadHTMLGlob("internal/adapter/http/web/templates/*")

	router.GET("/home", AuthMiddleware(handler.tokenSvc), handler.HomeGET)
}

func (ah *AuthHandler) HomeGET(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		_, msg := common.ErrorToHTTP(errors.New("internal server error"))
		redirectToLogin(c, msg)
		return
	}
	user, err := ah.auth.GetUserByID(userID.(domUser.UserID))
	if err != nil {
		_, msg := common.ErrorToHTTP(err)
		redirectToLogin(c, msg)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name":    user.Name,
		"surname": user.Surname,
		"role":    user.Role,
	})
}
