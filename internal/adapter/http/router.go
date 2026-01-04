package http

import (
	"time"

	"canteen-app/internal/adapter/http/api"
	"canteen-app/internal/adapter/http/common"
	"canteen-app/internal/adapter/http/web"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewRouter(authUC common.AuthUseCase, refreshTTL time.Duration) *gin.Engine {
	r := gin.Default()

	api.NewAuthHandler(r, authUC, refreshTTL)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	web.NewAuthHandler(r)

	return r
}
