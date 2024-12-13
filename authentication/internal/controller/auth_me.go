package controller

import (
	"context"
	"server/internal/entity"
	"server/pb"
	"server/util"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Me(ctx context.Context, req *pb.AuthMeRequest) (*pb.AuthMeResponse, error) {
	user := util.CurrentUser(ctx)
	data, err := m.authService.Me(ctx, &entity.AuthMe{
		UserId: user.UserId,
	})

	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.AuthMeResponse{}, status.Error(code, message)
	}

	return &pb.AuthMeResponse{
		Id:          data.UUID,
		Username:    data.Username,
		RoleCode:    data.RoleCode,
		RoleName:    data.RoleName,
		Permissions: data.Permissions,
	}, nil
}
