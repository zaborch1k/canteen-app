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
	refreshRepo := ram_storage.NewRefreshRepo()
	tokenSvc := jwtadapter.NewJWTTokenService([]byte("SECRET"), "issuer", []byte("SECRET2"), time.Hour*24*7)
	authUC := usecase.NewAuthUseCase(userRepo, tokenSvc, time.Duration(1)*time.Minute, refreshRepo)
	router := http.NewRouter(authUC, tokenSvc)

	return &App{
		router: router,
	}, nil
}

func (a App) Run(port string) error {
	return a.router.Run(port)
}
