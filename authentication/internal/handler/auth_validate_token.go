package handler

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) ValidateToken(ctx context.Context, req *pb.AuthValidateTokenRequest) (*pb.AuthValidateTokenResponse, error) {
	claim, err := m.authService.ValidateAccess(&ctx, &entity.AuthValidateAccess{
		Token: &req.Token,
	})

	if err.Errors != nil {
		return &pb.AuthValidateTokenResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}
	return &pb.AuthValidateTokenResponse{
		UserId:   claim.UserId,
		RoleCode: claim.RoleCode,
	}, nil
}
