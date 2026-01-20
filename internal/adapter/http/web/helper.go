package web

import (
	"log"
	"net/http"

	"canteen-app/internal/adapter/security/csrf"

	"github.com/gin-gonic/gin"
)

func redirectToAuthPage(c *gin.Context, path string, reason string) {
	c.SetCookie("flash_auth", reason, 60, "/", "", true, true)
	c.Redirect(http.StatusSeeOther, path)
	c.Abort()
}

func getFlash(c *gin.Context, name string) string {
	v, err := c.Cookie(name)
	if err != nil {
		return ""
	}
	c.SetCookie(name, "", -1, "/", "", true, true)
	return v
}

func setCsrfCookie(c *gin.Context) string {
	csrfToken, err := c.Cookie("csrf_token")
	if err != nil || csrfToken == "" {
		csrfToken, _ = csrf.NewToken()

		c.SetCookieData(&http.Cookie{
			Name:     "csrf_token",
			Value:    csrfToken,
			Path:     "/",
			Domain:   "",
			MaxAge:   0,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})
	}

	return csrfToken
}

func denyCSRF(c *gin.Context, reason string) {
	log.Println("csrf failed",
		"reason", reason,
		"ip", c.ClientIP(),
		"path", c.Request.URL.Path,
	)

	redirectToAuthPage(c, "/login", "session has expired")
}
