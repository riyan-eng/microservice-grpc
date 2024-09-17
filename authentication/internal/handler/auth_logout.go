package handler

import (
	"context"
	"server/internal/entity"
	"server/pb"
	"server/util"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Logout(ctx context.Context, req *pb.AuthLogoutRequest) (*pb.AuthLogoutResponse, error) {
	user := util.CurrentUser(ctx)
	err := m.authService.Logout(&ctx, &entity.AuthLogout{
		UserId: &user.UserId,
	})

	if err.Errors != nil {
		return &pb.AuthLogoutResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}

	return &pb.AuthLogoutResponse{}, nil
}
