package controller

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) Login(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	data, token, err := m.authService.Login(ctx, &entity.AuthLogin{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.AuthLoginResponse{}, status.Error(code, message)
	}

	return &pb.AuthLoginResponse{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		AccessExpired:  token.AccessExpired.String(),
		RefreshExpired: token.RefreshExpired.String(),
		Username:       data.Username,
	}, nil
}
