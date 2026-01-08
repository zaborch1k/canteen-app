package web

import (
	"canteen-app/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenService usecase.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil {
			log.Println(err.Error())
			redirectToAuthPage(c, "/login", "")
			return
		}

		claims, err := tokenService.ParseAccessToken(tokenStr)
		if err != nil {
			log.Println(err.Error())
			redirectToAuthPage(c, "/login", "")
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}
