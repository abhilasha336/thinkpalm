package controllers

import (
	"github.com/abhilasha336/thinkpalm/internal/dstructures"
	"github.com/abhilasha336/thinkpalm/internal/usecaseslogic"
	"github.com/gin-gonic/gin"
)

// ThinkpalmController struct holds router group and usecase inetrface
type ThinkpalmController struct {
	router  *gin.RouterGroup
	useCase usecaseslogic.ThinkpalmUsecaseImplements
	cfg     *dstructures.EnvConfig
}

// NewThinkpalmController used to pass value of router and usecases
func NewThinkpalmController(router *gin.RouterGroup, useCase usecaseslogic.ThinkpalmUsecaseImplements, cfg *dstructures.EnvConfig) *ThinkpalmController {
	return &ThinkpalmController{
		router:  router,
		useCase: useCase,
		cfg:     cfg,
	}
}

// InitRoutes function used to init all routes
func (think *ThinkpalmController) InitRoutes() {

	think.router.GET("/health", func(ctx *gin.Context) {
		think.HealthHandler(ctx)
	})
	think.router.POST("/register-user", func(ctx *gin.Context) {
		think.RegisterUser(ctx)
	})
	think.router.POST("/login-user", func(ctx *gin.Context) {
		think.LoginUser(ctx)
	})

	think.router.GET("/register", func(ctx *gin.Context) {
		think.Register(ctx)
	})
	think.router.GET("/login", func(ctx *gin.Context) {
		think.Login(ctx)
	})
	think.router.GET("/success", func(ctx *gin.Context) {
		think.Success(ctx)
	})
	think.router.GET("/failure", func(ctx *gin.Context) {
		think.Failure(ctx)
	})

}
