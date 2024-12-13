package router

import (
	"server/internal/api"
	rpcserver "server/rpc_server"

	"github.com/gin-gonic/gin"
)

type routeStruct struct {
	app     *gin.Engine
	handler *api.ServiceServer
}

func NewRouter(app *gin.Engine) *routeStruct {
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
