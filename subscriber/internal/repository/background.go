package repository

import (
	"context"
	"fmt"
	"server/internal/model"

	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
)

type BackgroundRepository interface {
	Create(ctx context.Context, message []byte, mdl *model.Background) error
}

type backgroundRepository struct {
	channel *amqp.Channel
	sqlxDB  *sqlx.DB
}

func (m *backgroundRepository) Create(ctx context.Context, message []byte, mdl *model.Background) error {
	if err := m.channel.Publish(
		"delayed_exchange",
		"notifications",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	); err != nil {
		return fmt.Errorf("13:%v", err)
	}

	sqlRslt, err := m.sqlxDB.NamedExecContext(ctx, `
	insert into notifications (id, status, type, body, error) values (:id, :status, :type, :body, :error)
	`, mdl)
	if err != nil {
		return fmt.Errorf("13:%v", err)
	}
	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return fmt.Errorf("13:%v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("5:no data entered")
	}

	return nil
}
