package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (think *ThinkpalmController) ClientRegisterForm(ctx *gin.Context) {
	// Pass the data to the template
	ctx.HTML(http.StatusOK, "client_register_form.html", nil)
}
