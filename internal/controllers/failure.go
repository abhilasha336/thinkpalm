package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (think *ThinkpalmController) Failure(ctx *gin.Context) {
	// Pass the data to the template
	ctx.HTML(http.StatusOK, "failure.html", nil)
}
