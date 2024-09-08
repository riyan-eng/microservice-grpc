package handler

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Login(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	data, token, err := m.authService.Login(&ctx, &entity.AuthLogin{
		Username: &req.Username,
		Password: &req.Password,
	})

	if err.Errors != nil {
		return &pb.AuthLoginResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}

	return &pb.AuthLoginResponse{
		AccessToken:    *token.AccessToken,
		RefreshToken:   *token.RefreshToken,
		AccessExpired:  token.AccessExpired.String(),
		RefreshExpired: token.RefreshExpired.String(),
		Username:       data.Username,
	}, nil
}
