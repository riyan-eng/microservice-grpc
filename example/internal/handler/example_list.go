package handler

import (
	"context"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) List(ctx context.Context, req *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	data, count, err := m.dao.NewExampleRepository().List(&ctx, &req.Search, &req.Limit, &req.Offset)
	if err.Errors != nil {
		return &pb.TaskListResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}

	dataRes := make([]*pb.TaskList, 0)
	for _, d := range *data {
		dataRes = append(dataRes, &pb.TaskList{
			Id:     d.Id,
			Name:   d.Name,
			Detail: d.Detail,
		})
	}

	return &pb.TaskListResponse{
		Data:      dataRes,
		TotalRows: *count,
	}, nil
}
