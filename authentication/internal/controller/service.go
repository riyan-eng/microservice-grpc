package controller

import (
	"server/internal/repository"
	"server/internal/service"
	"server/pb"
	"server/util"
)

type ServiceServer struct {
	pb.UnimplementedAuthServiceServer
	pb.UnimplementedPermissionServiceServer
	util        *util.Util
	dao         repository.DAO
	authService service.AuthService
}

func NewService(
	util *util.Util,
	dao repository.DAO,
	authService service.AuthService,
) *ServiceServer {
	return &ServiceServer{
		util:        util,
		dao:         dao,
		authService: authService,
	}
}
