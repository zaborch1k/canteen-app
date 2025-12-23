package http 

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/", func (c *gin.Context) {
        c.String(http.StatusOK, "hello")
    })

    return r
}
