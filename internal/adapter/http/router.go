package http

import (
	"time"

	docs "canteen-app/cmd/docs"
	"canteen-app/internal/adapter/http/api"
	"canteen-app/internal/adapter/http/common"
	"canteen-app/internal/usecase"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewRouter(authUC common.AuthUseCase, tokenService usecase.TokenService, refreshTTL time.Duration) *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	api.NewAuthHandler(r, authUC, tokenService, refreshTTL)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
