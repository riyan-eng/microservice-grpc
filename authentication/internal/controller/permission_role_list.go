package controller

import (
	"context"
	"server/pb"

	"google.golang.org/grpc/status"
)

func (m ServiceServer) RoleAccessList(ctx context.Context, req *pb.PermissionRoleAccessRequest) (*pb.PermissionRoleAccessListResponse, error) {
	data, err := m.dao.NewPermissionRepository().RoleAccessList(ctx)
	if err != nil {
		code, message := m.util.Error.ParsingPrefix(err)
		return &pb.PermissionRoleAccessListResponse{}, status.Error(code, message)
	}

	roles := make([]*pb.PermissionRoleAccessResponse, 0)
	for _, d := range data {
		permissions := make([]*pb.PermissionRoleAccessChildResponse, 0)
		for _, e := range d.Permissions {
			permissions = append(permissions, &pb.PermissionRoleAccessChildResponse{
				Id:         e.Id,
				Code:       e.Code,
				Name:       e.Name,
				Parent:     e.Parent,
				FullMethod: e.FullMethod,
			})
		}
		roles = append(roles, &pb.PermissionRoleAccessResponse{
			Name:        d.Name,
			Permissions: permissions,
		})
	}

	return &pb.PermissionRoleAccessListResponse{
		Data: roles,
	}, nil
}
