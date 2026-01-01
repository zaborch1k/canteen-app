package app

import (
	"time"

	"canteen-app/internal/adapter/http"
	jwtadapter "canteen-app/internal/adapter/jwt"
	"canteen-app/internal/adapter/repo/ram_storage"
	"canteen-app/internal/adapter/security"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func New() (*App, error) {
	userRepo := ram_storage.NewUserRepo()
	refreshRepo := ram_storage.NewRefreshRepo()

	accessTTL := time.Minute * 30
	refreshTTL := time.Hour * 24 * 30

	tokenSvc := jwtadapter.NewJWTTokenService([]byte("SECRET"), []byte("SECRET2"), accessTTL, refreshTTL, "issuer")
	bhasher := security.BcryptHasher{}
	authUC := usecase.NewAuthUseCase(userRepo, tokenSvc, refreshRepo, bhasher)
	router := http.NewRouter(authUC, refreshTTL)

	return &App{
		router: router,
	}, nil
}

func (a App) Run(port string) error {
	return a.router.Run(port)
}
