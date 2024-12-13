package repository

import (
	"consumer-trial/config"

	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
)

type DAO interface {
	NewNotificationRepository() NotificationRepository
}

type dao struct {
	SqlxDB  *sqlx.DB
	Channel *amqp.Channel
}

func NewDAO(conf *config.Config) DAO {
	return &dao{
		SqlxDB:  conf.SqlXDB,
		Channel: conf.Channel,
	}
}

func (m *dao) NewNotificationRepository() NotificationRepository {
	return &notificationRepository{
		sqlxDB:  m.SqlxDB,
		channel: m.Channel,
	}
}
