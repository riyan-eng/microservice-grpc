package handler

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Put(ctx context.Context, req *pb.TaskPutRequest) (*pb.TaskPutResponse, error) {
	err := m.exampleService.Put(&ctx, &entity.ExamplePut{
		Id:     &req.Id,
		Name:   &req.Name,
		Detail: &req.Detail,
	})

	if err.Errors != nil {
		return &pb.TaskPutResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}

	return &pb.TaskPutResponse{}, nil
}
