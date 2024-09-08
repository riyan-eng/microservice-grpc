package handler

import (
	"server/internal/repository"
	"server/internal/service"
	"server/pb"
)

type ServiceServer struct {
	pb.UnimplementedTaskServiceServer
	dao            repository.DAO
	exampleService service.ExampleService
}

func NewService(
	dao repository.DAO,
	exampleService service.ExampleService,
) *ServiceServer {
	return &ServiceServer{
		dao:            dao,
		exampleService: exampleService,
	}
}
