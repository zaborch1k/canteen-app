package http

import (
	"canteen-app/internal/adapter/http/api"
	domAuth "canteen-app/internal/domain/auth"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(authUC usecase.AuthUseCase, tokenService domAuth.TokenService) *gin.Engine {
	r := gin.Default()

	api.NewAuthHandler(r, authUC, tokenService)

	return r
}
