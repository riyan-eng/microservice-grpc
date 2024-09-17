package util

import "context"

type AccessTokenClaimsKey struct{}

func CurrentUser(ctx context.Context) AccessTokenClaims {
	return ctx.Value(AccessTokenClaimsKey{}).(AccessTokenClaims)
}
