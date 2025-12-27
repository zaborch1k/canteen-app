package app

import (
	"time"

	"canteen-app/internal/adapter/http"
	jwtadapter "canteen-app/internal/adapter/jwt"
	"canteen-app/internal/adapter/repo/ram_storage"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func New() (*App, error) {
	userRepo := ram_storage.NewUserRepo()
	tokenSvc := jwtadapter.NewJWTTokenService([]byte("SECRET"), "issuer")
	userUC := usecase.NewUserUseCase(userRepo, tokenSvc, time.Duration(30)*time.Minute)
	router := http.NewRouter(userUC, tokenSvc)

	return &App{
		router: router,
	}, nil
}

func (a App) Run(port string) error {
	return a.router.Run(port)
}
