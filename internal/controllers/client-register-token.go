package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (think *ThinkpalmController) ClientRegisterToken(ctx *gin.Context) {
	// Pass the data to the template
	ctx.HTML(http.StatusOK, "client_app.html", nil)
}
