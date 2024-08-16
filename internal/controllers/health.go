package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler function used too check health status of server
func (think *ThinkpalmController) HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"server-status": "available and health ok",
	})
}
