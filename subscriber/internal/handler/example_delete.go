package handler

import (
	"context"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Delete(ctx context.Context, req *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	err := m.dao.NewExampleRepository().Delete(ctx, req.Id)
	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.TaskDeleteResponse{}, status.Error(code, message)
	}

	return &pb.TaskDeleteResponse{}, nil
}
