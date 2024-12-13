package repository

import (
	"database/sql"
	"server/config"

	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type DAO interface {
	NewAuthRepository() AuthRepository
	NewPermissionRepository() PermissionRepository
}

type dao struct {
	SqlDB      *sql.DB
	SqlxDB     *sqlx.DB
	Cache      *redis.Client
	Permission *casbin.Enforcer
}

func NewDAO(conf *config.Config) DAO {
	return &dao{
		SqlDB:      conf.SqlDB,
		SqlxDB:     conf.SqlXDB,
		Cache:      conf.Cache,
		Permission: conf.Permission,
	}
}

func (m *dao) NewAuthRepository() AuthRepository {
	return &authRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
		cache:  m.Cache,
	}
}

func (m *dao) NewPermissionRepository() PermissionRepository {
	return &permissionRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
		cache:  m.Cache,
	}
}
