package controller

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) ValidatePermission(ctx context.Context, req *pb.AuthValidatePermissionRequest) (*pb.AuthValidatePermissionResponse, error) {
	err := m.authService.ValidatePermission(ctx, &entity.AuthValidatePermission{
		UserId:     req.UserId,
		FullMethod: req.FullMethod,
	})

	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.AuthValidatePermissionResponse{}, status.Error(code, message)
	}
	return &pb.AuthValidatePermissionResponse{}, nil
}
