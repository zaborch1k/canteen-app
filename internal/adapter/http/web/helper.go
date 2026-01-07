package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirectToLogin(c *gin.Context, reason string) {
	c.SetCookie("flash_auth", reason, 60, "/", "", true, true)
	c.Redirect(http.StatusSeeOther, "/register")
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
