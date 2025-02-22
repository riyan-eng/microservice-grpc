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
	authRpcServer, permissionRpcServer := rpcserver.AuthService()

	handler := api.NewService(
		// exampleService,
		// authService,
		exampleRpcServer,
		authRpcServer,
		permissionRpcServer,
	)
	return &routeStruct{
		app:     app,
		handler: handler,
	}
}
