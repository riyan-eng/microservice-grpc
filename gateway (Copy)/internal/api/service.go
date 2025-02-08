package api

import (
	"server/pb"
)

type ServiceServer struct {
	exampleRpcServer pb.TaskServiceClient
	authRpcServer    pb.AuthServiceClient
}

func NewService(
	exampleRpcServer pb.TaskServiceClient,
	authRpcServer pb.AuthServiceClient,
) *ServiceServer {
	return &ServiceServer{
		exampleRpcServer: exampleRpcServer,
		authRpcServer:    authRpcServer,
	}
}
