package controllers

import (
	"fmt"
	"net/http"

	"github.com/abhilasha336/thinkpalm/internal/dstructures"
	"github.com/gin-gonic/gin"
)

func (think *ThinkpalmController) LoginUser(ctx *gin.Context) {

	var loginReq dstructures.LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"details": "controller bind error" + err.Error()})
		return
	}
	fmt.Printf("%+v user struct", loginReq)

	if err := think.useCase.LoginUser(ctx, loginReq); err != nil {
		fmt.Println("controller error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"details": "user loginfailed"})

		return
	}
	ctx.JSON(http.StatusOK, gin.H{"details": "user registration successfull"})

}
