package http

import (
	"canteen-app/internal/adapter/http/api"
	"canteen-app/internal/adapter/http/common"
	"canteen-app/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouter(authUC common.AuthUseCase, tokenService usecase.TokenService, refreshTTL time.Duration) *gin.Engine {
	r := gin.Default()

	api.NewAuthHandler(r, authUC, tokenService, refreshTTL)

	return r
}
