package util

import (
	"context"
)

type AccessTokenClaimsKey struct{}

type AccessTokenClaims struct {
	UserId   string `json:"user_id"`
	RoleCode string `json:"role_code"`
}

func CurrentUser(ctx context.Context) AccessTokenClaims {
	return ctx.Value(AccessTokenClaimsKey{}).(AccessTokenClaims)
}
