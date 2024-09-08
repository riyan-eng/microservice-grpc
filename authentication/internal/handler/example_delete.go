package handler

import (
	"context"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Delete(ctx context.Context, req *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	err := m.dao.NewExampleRepository().Delete(&ctx, &req.Id)
	if err.Errors != nil {
		return &pb.TaskDeleteResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}

	return &pb.TaskDeleteResponse{}, nil
}
