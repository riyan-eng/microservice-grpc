package infrastructure

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMq() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	// defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("failed to open a channel: %v", err)
	}
	// defer ch.Close()

	// Declare the delayed exchange
	err = ch.ExchangeDeclare(
		"delayed_exchange",                     // name
		"x-delayed-message",                    // type
		true,                                   // durable
		false,                                  // auto-deleted
		false,                                  // internal
		false,                                  // no-wait
		amqp.Table{"x-delayed-type": "direct"}, // arguments
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to declare exchange: %v", err)
	}

	log.Println("rabbitmq ch: connection established")
	return conn, ch, nil
}
