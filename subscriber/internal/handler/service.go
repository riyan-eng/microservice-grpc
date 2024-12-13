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
	pb.UnimplementedBackgroundServiceServer
	dao               repository.DAO
	exampleService    service.ExampleService
	backgroundService service.BackgroundService
}

func NewService(
	util *util.Util,
	dao repository.DAO,
	exampleService service.ExampleService,
	backgroundService service.BackgroundService,
) *ServiceServer {
	return &ServiceServer{
		util:              util,
		dao:               dao,
		exampleService:    exampleService,
		backgroundService: backgroundService,
	}
}
