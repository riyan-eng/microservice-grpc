package controller

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) ValidateToken(ctx context.Context, req *pb.AuthValidateTokenRequest) (*pb.AuthValidateTokenResponse, error) {
	claim, err := m.authService.ValidateAccess(ctx, &entity.AuthValidateAccess{
		Token: req.Token,
	})

	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.AuthValidateTokenResponse{}, status.Error(code, message)
	}
	return &pb.AuthValidateTokenResponse{
		UserId:   claim.UserId,
		RoleCode: claim.RoleCode,
	}, nil
}
