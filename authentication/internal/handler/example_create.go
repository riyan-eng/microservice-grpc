package handler

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Create(ctx context.Context, req *pb.TaskCreateRequest) (*pb.TaskCreateResponse, error) {
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
