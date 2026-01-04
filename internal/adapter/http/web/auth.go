package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthHandler(router *gin.Engine) {
	router.LoadHTMLGlob("internal/adapter/http/web/templates/*")

	{
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"name": "sdflsdflsdkfl",
			})
		})

		router.POST("/", func(c *gin.Context) {
			input := c.PostForm("input")
			fmt.Println(input)
		})
	}
}
