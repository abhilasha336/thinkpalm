package controllers

import (
	"net/http"

	"github.com/abhilasha336/thinkpalm/internal/dstructures"
	"github.com/gin-gonic/gin"
)

func (think *ThinkpalmController) RegisterUser(ctx *gin.Context) {

	var loginReq dstructures.LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"details": "controller bind error" + err.Error()})
		return
	}
	if err := think.useCase.RegisterUser(ctx, loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"details": "user registration failed"})

		return
	}
	ctx.JSON(http.StatusOK, gin.H{"details": "user registration successfull"})

}
