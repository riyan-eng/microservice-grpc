package handler

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"github.com/samborkent/uuidv7"
	"google.golang.org/grpc/status"
)

func (m ServiceServer) BackgroundCreate(ctx context.Context, req *pb.BackgroundCreateRequest) (*pb.BackgroundCreateResponse, error) {
	if err := m.backgroundService.Create(ctx, &entity.BackgroundCreate{
		Id:   uuidv7.New().String(),
		Type: req.Type,
		Body: entity.BackgroundCreateBody{
			To:          req.Body.To,
			Name:        req.Body.Name,
			Description: req.Body.Desc,
			Var1:        req.Body.V1,
			Var2:        req.Body.V2,
			Var3:        req.Body.V3,
			Var4:        req.Body.V4,
			Var5:        req.Body.V5,
		},
	}); err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.BackgroundCreateResponse{}, status.Error(code, message)
	}
	return &pb.BackgroundCreateResponse{}, nil
}
