package repository

import (
	"database/sql"
	"server/config"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	NewExampleRepository() ExampleRepository
}

type dao struct {
	SqlDB  *sql.DB
	SqlxDB *sqlx.DB
}

func NewDAO(conf *config.Config) DAO {
	return &dao{
		SqlDB:  conf.SqlDB,
		SqlxDB: conf.SqlXDB,
	}
}

func (m *dao) NewExampleRepository() ExampleRepository {
	return &exampleRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
	}
}
