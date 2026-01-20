package http

import (
	"time"

	"canteen-app/internal/adapter/http/api"
	"canteen-app/internal/adapter/http/common"
	"canteen-app/internal/adapter/http/web"
	"canteen-app/internal/usecase"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	authUC common.AuthUseCase,
	accessTTL time.Duration,
	refreshTTL time.Duration,
	tokenSvc usecase.TokenService,
	validator Validator,
) *gin.Engine {
	r := gin.Default()

	api.NewAuthHandler(r, authUC, refreshTTL, validator)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	web.NewAuthHandler(r, authUC, accessTTL, refreshTTL, tokenSvc, validator)

	return r
}
