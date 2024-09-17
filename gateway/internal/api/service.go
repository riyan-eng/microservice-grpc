package api

import (
	"server/internal/service"
	"server/pb"
)

type ServiceServer struct {
	// authService      service.AuthService
	perangkatService service.PerangkatService
	objectService    service.ObjectService
	exampleRpcServer pb.TaskServiceClient
	authRpcServer    pb.AuthServiceClient
}

func NewService(
	// authService service.AuthService,
	perangkatService service.PerangkatService,
	objectService service.ObjectService,
	exampleRpcServer pb.TaskServiceClient,
	authRpcServer pb.AuthServiceClient,
) *ServiceServer {
	return &ServiceServer{
		// authService:      authService,
		perangkatService: perangkatService,
		objectService:    objectService,
		exampleRpcServer: exampleRpcServer,
		authRpcServer:    authRpcServer,
	}
}
