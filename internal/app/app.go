package app

import (
    "canteen-app/internal/adapter/http"

    "github.com/gin-gonic/gin"
)

type App struct {
    router *gin.Engine;
}

func New() (*App, error) {
    router := http.NewRouter() 

    return &App {
        router: router,
    }, nil
}

func (a App) Run(port string) error {
    return a.router.Run(port)
}
