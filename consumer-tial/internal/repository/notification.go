package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/time/rate"
)

var limiterWhatsApp = rate.NewLimiter(1, 1)

type NotificationRepository interface {
	WhatsApp(ctx context.Context, id, status string) error
}

type notificationRepository struct {
	channel *amqp.Channel
	sqlxDB  *sqlx.DB
}

func (m *notificationRepository) WhatsApp(ctx context.Context, id, status string) error {
	if err := limiterWhatsApp.Wait(context.TODO()); err != nil {
		return fmt.Errorf("rate limiter error: %v", err)
	}

	switch status {
	case "PROCESS":
		log.Println("hot wa api")
		_, err := m.sqlxDB.ExecContext(ctx, `update notifications set status='SUCCESS' where id=$1`, id)
		if err != nil {
			return err
		}
	default:
		_, err := m.sqlxDB.ExecContext(ctx, `update notifications set status='FAILED', error=$2 where id=$1`, id, fmt.Errorf("api error").Error())
		if err != nil {
			return err
		}
	}
	return nil
}
