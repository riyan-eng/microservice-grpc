package interceptor

import (
	"context"
	"fmt"
	"server/pb"
	rpcserver "server/rpc_server"
	"server/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		fmt.Println("--> unary interceptor")
		ctx, err := Auth(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func Auth(ctx context.Context, method string) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	if methodAuthWhiteList[method] {
		return ctx, nil
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	// hit auth service
	claim, err := rpcserver.AuthService().ValidateToken(ctx, &pb.AuthValidateTokenRequest{
		Token: values[0],
	})

	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "token doesn't valid")
	}

	if !methodPermWhiteList[method] {
		_, err := rpcserver.AuthService().ValidatePermission(ctx, &pb.AuthValidatePermissionRequest{
			UserId:     claim.UserId,
			FullMethod: method,
		})
		if err != nil {
			return ctx, status.Errorf(codes.PermissionDenied, "you are not allowed")
		}
	}

	newClaim := util.AccessTokenClaims{
		UserId:   claim.UserId,
		RoleCode: claim.RoleCode,
	}

	ctx = context.WithValue(ctx, util.AccessTokenClaimsKey{}, newClaim)

	return ctx, nil
}

var methodAuthWhiteList = map[string]bool{
	"/proto.TaskService/List":   true,
	"/proto.TaskService/Detail": true,
}

var methodPermWhiteList = map[string]bool{
	// "/proto.TaskService/Create": true,
}
