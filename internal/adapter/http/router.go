package http

import (
	"canteen-app/internal/adapter/http/api"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(userUC usecase.UserUseCase) *gin.Engine {
	r := gin.Default()

	api.NewAuthHandler(r, userUC)

	return r
}
