package controller

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Refresh(ctx context.Context, req *pb.AuthRefreshRequest) (*pb.AuthRefreshResponse, error) {
	token, err := m.authService.Refresh(ctx, &entity.AuthRefresh{
		Token: req.Token,
	})

	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.AuthRefreshResponse{}, status.Error(code, message)
	}

	return &pb.AuthRefreshResponse{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		AccessExpired:  token.AccessExpired.String(),
		RefreshExpired: token.RefreshExpired.String(),
	}, nil
}
