package rabbitmq

import (
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

// PublishWithDelay sends a message to the delayed exchange with a delay.
func PublishWithDelay(ch *amqp.Channel, message string, delayMs int, retryCount int) error {
	headers := amqp.Table{
		"x-delay":     delayMs,
		"retry-count": strconv.Itoa(retryCount),
	}
	err := ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:     headers,
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}
	// log.Printf("Published message: %s with delay: %d ms (retry count: %d)", message, delayMs, retryCount)
	return nil
}
