package handler

import (
	"context"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Detail(ctx context.Context, req *pb.TaskDetailRequest) (*pb.TaskDetailResponse, error) {
	data, err := m.dao.NewExampleRepository().Detail(ctx, req.Id)
	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.TaskDetailResponse{}, status.Error(code, message)
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
