package http

import (
	"canteen-app/internal/adapter/http/api"
	domAuth "canteen-app/internal/domain/auth"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(userUC usecase.UserUseCase, tokenService domAuth.TokenService) *gin.Engine {
	r := gin.Default()

	api.NewAuthHandler(r, userUC, tokenService)

	return r
}
