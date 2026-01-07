package api

import (
	"canteen-app/internal/adapter/http/common"

	"github.com/gin-gonic/gin"
)

func writeError(c *gin.Context, err error) {
	status, msg := common.ErrorToHTTP(err)
	c.JSON(status, gin.H{"error": msg})
}
