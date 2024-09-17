package handler

import (
	"context"
	"server/internal/entity"
	"server/pb"
	"server/util"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Me(ctx context.Context, req *pb.AuthMeRequest) (*pb.AuthMeResponse, error) {
	user := util.CurrentUser(ctx)
	data, err := m.authService.Me(&ctx, &entity.AuthMe{
		UserId: &user.UserId,
	})

	if err.Errors != nil {
		return &pb.AuthMeResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}

	return &pb.AuthMeResponse{
		Username: data.Username,
		RoleCode: data.RoleCode,
		RoleName: data.RoleName,
	}, nil
}
