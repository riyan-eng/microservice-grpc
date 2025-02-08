package main

import (
	"fmt"
	"runtime"
	"server/config"
	"server/env"
	"server/infrastructure"
	"server/internal/router"
	"server/middleware"

	_ "server/docs"

	"github.com/gofiber/fiber/v3"

	// hertz-swagger middleware
	// swagger embed files
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/mvrilo/go-redoc"
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
	app := fiber.New()

	// documentation
	docsss := redoc.Redoc{
		Title:       "Sisalak",
		Description: "Sisalak API Description",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/try/doc.json",
		DocsPath:    "/documentation",
	}
	app.Use(middleware.Documentation(docsss))

	// middleware
	app.Use(recoverer.New())
	app.Use(middleware.RequestId())
	app.Use(middleware.Limiter())
	app.Use(infrastructure.LocalizerMiddleware())
	app.Use(middleware.Cors())

	// router
	routers := router.NewRouter(app)
	routers.Index()
	routers.Authentication()
	routers.Example()

	// startup log
	fmt.Println("server run on:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)

	app.Listen(env.NewEnv().SERVER_HOST + ":" + env.NewEnv().SERVER_PORT)
}
