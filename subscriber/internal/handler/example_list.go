package handler

import (
	"context"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) List(ctx context.Context, req *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	data, count, err := m.dao.NewExampleRepository().List(ctx, req.Search, req.Limit, req.Offset)
	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.TaskListResponse{}, status.Error(code, message)
	}

	dataRes := make([]*pb.TaskList, 0)
	for _, d := range data {
		dataRes = append(dataRes, &pb.TaskList{
			Id:     d.Id,
			Name:   d.Name,
			Detail: d.Detail,
		})
	}

	return &pb.TaskListResponse{
		Data:      dataRes,
		TotalRows: count,
	}, nil
}
