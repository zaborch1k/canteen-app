package web

import (
	"canteen-app/internal/adapter/security/csrf"
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

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieToken, err := c.Cookie("csrf_token")
		if err != nil || cookieToken == "" {
			denyCSRF(c, "missing_csrf_cookie")
			return
		}

		formToken := c.PostForm("csrf_token")
		if formToken == "" {
			denyCSRF(c, "missing_csrf_form")
			return
		}

		if !csrf.Compare(cookieToken, formToken) {
			denyCSRF(c, "mismatch")
			return
		}

		c.Next()
	}
}
