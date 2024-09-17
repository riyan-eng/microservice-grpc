package main

import (
	"fmt"
	"runtime"
	"server/config"
	"server/env"
	"server/infrastructure"
	"server/internal/repository"
	"server/internal/router"
	"server/middleware"

	_ "server/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// hertz-swagger middleware
	// swagger embed files
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU - 1)
	}
	env.LoadEnvironmentFile()
	env.NewEnv()

	config.NewLimiterStore()
	config.NewLogger()

	infrastructure.ConnectSqlDB()
	infrastructure.ConnectSqlxDB()
	infrastructure.ConnRedis()
	infrastructure.NewLocalizer()
}

// @title Sisalak
// @description This is a Sisalak Api Documentation.

// @contact.name hertz-contrib
// @contact.url https://github.com/hertz-contrib

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// create instance
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	// documentation
	app.GET("/try/*any", swagger.WrapHandler(swaggerFiles.Handler))
	app.Use(ginredoc.New(redoc.Redoc{
		Title:       "Sisalak",
		Description: "Sisalak API Description",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/try/doc.json",
		DocsPath:    "/documentation",
	}))

	// middleware
	app.Use(gin.Recovery())
	app.Use(middleware.RequestId())
	app.Use(middleware.Logger())
	app.Use(middleware.Limiter())
	app.Use(infrastructure.LocalizerMiddleware())

	corsconfig := cors.DefaultConfig()
	corsconfig.AllowCredentials = true
	corsconfig.AllowOrigins = []string{"*"}
	corsconfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	app.Use(cors.New(corsconfig))

	// service
	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.SqlxDB, infrastructure.Redis)

	// router
	routers := router.NewRouter(app, &dao)
	routers.Index()
	routers.Authentication()
	routers.Example()
	routers.Perangkat()
	routers.Object()
	routers.Export()

	// startup log
	fmt.Println("server run on:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)

	app.Run(env.NewEnv().SERVER_HOST + ":" + env.NewEnv().SERVER_PORT)
}
