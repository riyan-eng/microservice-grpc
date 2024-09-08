package handler

import (
	"context"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Detail(ctx context.Context, req *pb.TaskDetailRequest) (*pb.TaskDetailResponse, error) {
	data, err := m.dao.NewExampleRepository().Detail(&ctx, &req.Id)
	if err.Errors != nil {
		return &pb.TaskDetailResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}

	dataRes := &pb.TaskDetail{
		Id:     data.Id,
		Name:   data.Name,
		Detail: data.Detail,
	}

	return &pb.TaskDetailResponse{
		Data: dataRes,
	}, nil
}
