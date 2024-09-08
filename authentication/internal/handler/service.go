package handler

import (
	"server/internal/repository"
	"server/internal/service"
	"server/pb"
)

type ServiceServer struct {
	pb.UnimplementedTaskServiceServer
	pb.UnimplementedAuthServiceServer
	dao            repository.DAO
	exampleService service.ExampleService
	authService    service.AuthService
}

func NewService(
	dao repository.DAO,
	exampleService service.ExampleService,
	authService service.AuthService,
) *ServiceServer {
	return &ServiceServer{
		dao:            dao,
		exampleService: exampleService,
		authService:    authService,
	}
}
