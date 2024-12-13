package handler

import (
	"server/internal/repository"
	"server/internal/service"
	"server/pb"
	"server/util"
)

type ServiceServer struct {
	util *util.Util
	pb.UnimplementedTaskServiceServer
	dao            repository.DAO
	exampleService service.ExampleService
}

func NewService(
	util *util.Util,
	dao repository.DAO,
	exampleService service.ExampleService,
) *ServiceServer {
	return &ServiceServer{
		util:           util,
		dao:            dao,
		exampleService: exampleService,
	}
}
