package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (think *ThinkpalmController) Register(ctx *gin.Context) {

	var loginReq LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"details": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"details": loginReq})

}
