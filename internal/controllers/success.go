package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (think *ThinkpalmController) Success(ctx *gin.Context) {
	// Pass the data to the template
	ctx.HTML(http.StatusOK, "success.html", nil)
}
