package http

import (
	"canteen-app/internal/adapter/http/api"
	"canteen-app/internal/adapter/http/common"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(authUC common.AuthUseCase, tokenService usecase.TokenService) *gin.Engine {
	r := gin.Default()

	api.NewAuthHandler(r, authUC, tokenService)

	return r
}
