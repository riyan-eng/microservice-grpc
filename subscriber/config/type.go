package config

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	SqlDB   *sql.DB
	SqlXDB  *sqlx.DB
	Channel *amqp.Channel
}
