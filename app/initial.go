package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abhilasha336/thinkpalm/configenv"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/tuneverse/toolkit/middleware"
)

// set up all initial configurations for server
func InitialRun() {
	// init the env config
	cfg, err := configenv.LoadConfig(consts.AppName)
	if err != nil {
		panic(err)
	}

	// database connection
	pgsqlDB, err := driver.ConnectDB(cfg.Db)
	if err != nil {
		log.Fatalf("unable to connect the database")
		return
	}

	// here initalizing the router
	router := initRouter()
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	api := router.Group("/api")

	// m := middlewares.NewMiddlewares(cfg, pgsqlDB)
	// api.Use(m.JwtMiddleware())

	// middleware from toolkit/middleware  to intialise log middleware
	api.Use(middleware.LogMiddleware(map[string]interface{}{}))

	//middleware from toolkit/middleware  to intialise version middleware
	api.Use(middleware.APIVersionGuard(middleware.VersionOptions{
		AcceptedVersions: cfg.AcceptedVersions,
	}))

	// complete user related initialization
	{

		// repo initialization
		oauthRepo := repo.NewOauthRepo(pgsqlDB, cfg)
		// initilizing usecases
		oauthUseCases := usecases.NewOauthUseCase(oauthRepo, entities.OAuthData{})
		// initalizing controllers
		oauthControllers := controllers.NewOauthController(api, oauthUseCases, cfg, &entities.OAuthCredentials{})
		// init the routes
		oauthControllers.InitRoutes()

	}

	// run the app
	launch(cfg, router)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.LoadHTMLGlob("templates/*")

	// CORS
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "client_id", "client_secret", "partner_id", "provider"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// },
		MaxAge: 12 * time.Hour,
	}))

	// common middlewares should be added here

	return router
}

// launch
func launch(cfg *entities.EnvConfig, router *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.Port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Println("Server listening in...", cfg.Port)
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1) // kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}
