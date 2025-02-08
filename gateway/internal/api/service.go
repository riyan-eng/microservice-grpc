package api

import (
	"server/pb"
)

type ServiceServer struct {
	exampleRpcServer    pb.TaskServiceClient
	authRpcServer       pb.AuthServiceClient
	permissionRpcServer pb.PermissionServiceClient
}

func NewService(
	exampleRpcServer pb.TaskServiceClient,
	authRpcServer pb.AuthServiceClient,
	permissionRpcServer pb.PermissionServiceClient,
) *ServiceServer {
	return &ServiceServer{
		exampleRpcServer:    exampleRpcServer,
		authRpcServer:       authRpcServer,
		permissionRpcServer: permissionRpcServer,
	}
}
