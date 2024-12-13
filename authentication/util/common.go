package util

import (
	"context"

	"github.com/bytedance/sonic"
)

type AccessTokenClaimsKey struct{}

func CurrentUser(ctx context.Context) AccessTokenClaims {
	return ctx.Value(AccessTokenClaimsKey{}).(AccessTokenClaims)
}

func UnmarshalConverter[T any](s string) (data T) {
	sonic.UnmarshalString(s, &data)
	return
}
