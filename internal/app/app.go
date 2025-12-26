package app

import (
	"canteen-app/internal/adapter/http"
	"canteen-app/internal/adapter/repo/ram_storage"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func New() (*App, error) {
	userRepo := ram_storage.NewUserRepo()
	userUC := usecase.NewUserUseCase(userRepo)
	router := http.NewRouter(userUC)

	return &App{
		router: router,
	}, nil
}

func (a App) Run(port string) error {
	return a.router.Run(port)
}
