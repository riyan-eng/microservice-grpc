package handler

import (
	"context"
	"fmt"
	"server/internal/entity"
	"server/pb"
	"server/util"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Create(ctx context.Context, req *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
	user := util.CurrentUser(ctx)
	fmt.Println(user.UserId)
	fmt.Println(user.RoleCode)

	err := m.exampleService.Create(&ctx, &entity.ExampleCreate{
		Id:     &req.Id,
		Name:   &req.Name,
		Detail: &req.Detail,
	})

	if err.Errors != nil {
		return &pb.TaskCreateResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}

	return &pb.TaskCreateResponse{}, nil
}
