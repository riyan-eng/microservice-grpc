package handler

import (
	"context"
	"server/internal/entity"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) ValidatePermission(ctx context.Context, req *pb.AuthValidatePermissionRequest) (*pb.AuthValidatePermissionResponse, error) {
	err := m.authService.ValidatePermission(&ctx, &entity.AuthValidatePermission{
		UserId:     &req.UserId,
		FullMethod: &req.FullMethod,
	})

	if err.Errors != nil {
		return &pb.AuthValidatePermissionResponse{}, status.Error(err.StatusCode, err.Errors.Error())
	}
	return &pb.AuthValidatePermissionResponse{}, nil
}
