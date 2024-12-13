package repository

import (
	"database/sql"
	"server/config"

	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
)

type DAO interface {
	NewExampleRepository() ExampleRepository
	NewBackgroundRepository() BackgroundRepository
}

type dao struct {
	SqlDB   *sql.DB
	SqlxDB  *sqlx.DB
	Channel *amqp.Channel
}

func NewDAO(conf *config.Config) DAO {
	return &dao{
		SqlDB:   conf.SqlDB,
		SqlxDB:  conf.SqlXDB,
		Channel: conf.Channel,
	}
}

func (m *dao) NewExampleRepository() ExampleRepository {
	return &exampleRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
	}
}

func (m *dao) NewBackgroundRepository() BackgroundRepository {
	return &backgroundRepository{
		channel: m.Channel,
		sqlxDB:  m.SqlxDB,
	}
}
