package controller

import (
	"context"
	"server/internal/entity"
	"server/pb"
	"server/util"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Logout(ctx context.Context, req *pb.AuthLogoutRequest) (*pb.AuthLogoutResponse, error) {
	user := util.CurrentUser(ctx)
	err := m.authService.Logout(ctx, &entity.AuthLogout{
		UserId: user.UserId,
	})

	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.AuthLogoutResponse{}, status.Error(code, message)
	}

	return &pb.AuthLogoutResponse{}, nil
}
