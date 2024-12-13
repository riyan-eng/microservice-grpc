package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Util struct {
	Token    *Token
	Security *Security
	Error    *Error
}

type Security struct{}

type Error struct{}

type AccessTokenClaims struct {
	UserId   string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UUID     string `json:"id"`
	jwt.RegisteredClaims
}

type AccessTokenCached struct {
	AccessUID string `json:"access"`
}

type RefreshTokenClaims struct {
	UserId   string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UUID     string `json:"id"`
	jwt.RegisteredClaims
}

type RefreshTokenCached struct {
	RefreshUID string `json:"refresh"`
}

type ResetTokenClaims struct {
	UserId string `json:"user_id"`
	UUID   string `json:"id"`
	jwt.RegisteredClaims
}

type ResetTokenCached struct {
	ResetUID string `json:"reset"`
}

type Token struct {
	Cache *redis.Client
}
