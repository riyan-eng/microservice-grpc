package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	NewExampleRepository() ExampleRepository
}

type dao struct {
	SqlDB  *sql.DB
	SqlxDB *sqlx.DB
}

func NewDAO(sqlDB *sql.DB, sqlxDB *sqlx.DB) DAO {
	return &dao{
		SqlDB:  sqlDB,
		SqlxDB: sqlxDB,
	}
}

func (m *dao) NewExampleRepository() ExampleRepository {
	return &exampleRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
	}
}
