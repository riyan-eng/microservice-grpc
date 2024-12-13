package config

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	SqlDB  *sql.DB
	SqlXDB *sqlx.DB
}
