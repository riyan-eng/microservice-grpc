package util

import (
	"context"
	"encoding/json"
	"fmt"
	"server/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

var mySigningKey = []byte("AllYourBase")

func NewToken(cache *redis.Client) *Token {
	return &Token{
		Cache: cache,
	}
}

func (t *Token) CreateAccess(ctx context.Context, userId, userRole string) (string, *jwt.NumericDate, error) {
	expired := new(jwt.NumericDate)

	expiredTime := time.Minute * env.NewEnv().JWT_EXPIRED_ACCESS

	tokenUUID := uuid.NewString()
	claims := AccessTokenClaims{
		userId,
		userRole,
		tokenUUID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return tokenString, expired, fmt.Errorf("can't signed the token")
	}

	cachedJson, err := json.Marshal(AccessTokenCached{
		AccessUID: claims.UUID,
	})
	if err != nil {
		return tokenString, expired, fmt.Errorf("can't marshal access token")
	}

	if err := t.Cache.Set(ctx, fmt.Sprintf("access-token-%s", userId), string(cachedJson), expiredTime).Err(); err != nil {
		return tokenString, expired, fmt.Errorf("can't cached access token")
	}
	return tokenString, claims.ExpiresAt, nil
}

func (t *Token) ParseAccess(tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*AccessTokenClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

func (t *Token) ValidateAccess(ctx context.Context, claims *AccessTokenClaims) error {
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, err := t.Cache.Get(ctx, fmt.Sprintf("access-token-%s", claims.UserId)).Result()
		if err != nil {
			return fmt.Errorf("token not found")
		}
		cachedTokens := new(AccessTokenCached)
		err = json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string = cachedTokens.AccessUID
		if err != nil || tokenUID != claims.UUID {
			return fmt.Errorf("token not found")
		}
		// if err != nil {
		// 	return fmt.Errorf("token not found")
		// }
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (t *Token) CreateRefresh(ctx context.Context, userId, userRole string) (string, *jwt.NumericDate, error) {
	expired := new(jwt.NumericDate)

	expiredTime := time.Minute * env.NewEnv().JWT_EXPIRED_REFRESH

	tokenUUID := uuid.NewString()
	claims := RefreshTokenClaims{
		userId,
		userRole,
		tokenUUID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return tokenString, expired, fmt.Errorf("can't signed the token")
	}

	cachedJson, err := json.Marshal(RefreshTokenCached{
		RefreshUID: claims.UUID,
	})
	if err != nil {
		return tokenString, expired, fmt.Errorf("can't marshal refresh token")
	}

	if err := t.Cache.Set(ctx, fmt.Sprintf("refresh-token-%s", userId), string(cachedJson), expiredTime).Err(); err != nil {
		return tokenString, expired, fmt.Errorf("can't cached refresh token")
	}
	return tokenString, claims.ExpiresAt, nil
}

func (t *Token) ParseRefresh(tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*RefreshTokenClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

func (t *Token) ValidateRefresh(ctx context.Context, claims *RefreshTokenClaims) error {
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, err := t.Cache.Get(ctx, fmt.Sprintf("refresh-token-%s", claims.UserId)).Result()
		if err != nil {
			return fmt.Errorf("token not found")
		}
		cachedTokens := new(RefreshTokenCached)
		err = json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string = cachedTokens.RefreshUID
		if err != nil || tokenUID != claims.UUID {
			return fmt.Errorf("token not found")
		}
		// if err != nil {
		// 	return fmt.Errorf("token not found")
		// }
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (t *Token) CreateReset(ctx context.Context, userId string) (string, *jwt.NumericDate, error) {
	expired := new(jwt.NumericDate)

	expiredTime := time.Minute * env.NewEnv().JWT_EXPIRED_RESET

	tokenUUID := uuid.NewString()
	claims := ResetTokenClaims{
		userId,
		tokenUUID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return tokenString, expired, fmt.Errorf("can't signed the token")
	}

	cachedJson, err := json.Marshal(ResetTokenCached{
		ResetUID: claims.UUID,
	})
	if err != nil {
		return tokenString, expired, fmt.Errorf("can't marshal reset token")
	}

	if err := t.Cache.Set(ctx, fmt.Sprintf("reset-token-%s", userId), string(cachedJson), expiredTime).Err(); err != nil {
		return tokenString, expired, fmt.Errorf("can't cached reset token")
	}
	return tokenString, claims.ExpiresAt, nil
}

func (t *Token) ParseReset(tokenString string) (*ResetTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ResetTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*ResetTokenClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

func (t *Token) ValidateReset(ctx context.Context, claims *ResetTokenClaims) error {
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, err := t.Cache.Get(ctx, fmt.Sprintf("reset-token-%s", claims.UserId)).Result()
		if err != nil {
			return fmt.Errorf("token not found")
		}
		cachedTokens := new(ResetTokenCached)
		err = json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string = cachedTokens.ResetUID
		if err != nil || tokenUID != claims.UUID {
			return fmt.Errorf("token not found")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
