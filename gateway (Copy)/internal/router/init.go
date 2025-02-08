package router

import (
	"server/internal/api"
	rpcserver "server/rpc_server"

	"github.com/gofiber/fiber/v3"
)

type routeStruct struct {
	app     *fiber.App
	handler *api.ServiceServer
}

func NewRouter(app *fiber.App) *routeStruct {
	// exampleService := service.NewExampleService(dao)
	// authService := service.NewAuthService(dao)

	exampleRpcServer := rpcserver.ExampleService()
	authRpcServer := rpcserver.AuthService()

	handler := api.NewService(
		// exampleService,
		// authService,
		exampleRpcServer,
		authRpcServer,
	)
	return &routeStruct{
		app:     app,
		handler: handler,
	}
}
