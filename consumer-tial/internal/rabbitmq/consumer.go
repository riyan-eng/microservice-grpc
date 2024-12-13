package rabbitmq

import (
	"consumer-trial/config"
	"consumer-trial/internal/repository"
	"consumer-trial/internal/route"
	"consumer-trial/internal/service"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
)

const PROCESS string = "PROCESS"
const DISCARD string = "DISCARD"

// ConsumeAndProcessMessages consumes messages from the queue and processes them.
func ConsumeAndProcessMessages(sqlxDB *sqlx.DB, ch *amqp.Channel) {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	myConfig := config.Config{
		SqlXDB:  sqlxDB,
		Channel: ch,
	}
	dao := repository.NewDAO(&myConfig)
	serviceNotification := service.NewNotification(&dao)
	routing := route.NewRoute(
		serviceNotification,
	)

	for msg := range msgs {
		message := string(msg.Body)
		retryCount := 0
		if count, ok := msg.Headers["retry-count"]; ok {
			retryCount, _ = strconv.Atoi(count.(string))
		}

		if err := routing.Route(PROCESS, message); err != nil {
			if retryCount < maxRetries {
				// log.Printf("Message failed: %s. Retrying...", message)
				retryCount++
				newDelay := initialDelayMs * (retryDelayFactor * retryCount)
				_ = PublishWithDelay(ch, message, newDelay, retryCount)
			} else {
				routing.Route(DISCARD, message)
				// log.Printf("Max retries reached for message: %s. Discarding.", message)
			}
		} else {
			// log.Printf("Successfully processed message: %s", message)
		}
	}
}
