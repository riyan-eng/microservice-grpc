package config

import (
	"database/sql"

	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	SqlDB      *sql.DB
	SqlXDB     *sqlx.DB
	Cache      *redis.Client
	Permission *casbin.Enforcer
}
