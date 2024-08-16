package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (think *ThinkpalmController) LogIn(ctx *gin.Context) {

	// Pass the data to the template
	ctx.HTML(http.StatusOK, "login.html", nil)

}
