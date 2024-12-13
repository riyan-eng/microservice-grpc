package interceptor

import (
	"context"
	"fmt"
	"server/internal/controller"
	"server/pb"
	"server/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type unaryInterceptor struct {
	service *controller.ServiceServer
}

func NewUnaryInterceptor(service *controller.ServiceServer) *unaryInterceptor {
	return &unaryInterceptor{
		service: service,
	}
}

func (m *unaryInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		fmt.Println("--> unary interceptor")
		ctx, err := m.Auth(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (m *unaryInterceptor) Auth(ctx context.Context, method string) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	if methodAuthWhiteList[method] {
		fmt.Println("ini non auth")
		return ctx, nil
	}

	fmt.Println("ini auth")
	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	// hit auth service
	claim, err := m.service.ValidateToken(ctx, &pb.AuthValidateTokenRequest{
		Token: values[0],
	})
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "token doesn't valid")
	}

	newClaim := util.AccessTokenClaims{
		UserId:   claim.UserId,
		RoleCode: claim.RoleCode,
	}

	ctx = context.WithValue(ctx, util.AccessTokenClaimsKey{}, newClaim)
	return ctx, nil
}

var methodAuthWhiteList = map[string]bool{
	"/proto.AuthService/Login":                true,
	"/proto.AuthService/Refresh":              true,
	"/proto.AuthService/ValidatePermission":   true,
	"/proto.AuthService/ValidateToken":        true,
	"/proto.PermissionService/RoleAccessList": true,
}
