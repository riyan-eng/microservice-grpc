package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const rabbitMQURL = "amqp://admin:admin@localhost:5672/"

// Connect establishes a connection to RabbitMQ and opens a channel.
func Connect() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	log.Println("rabbitmq ch: connection established")
	return conn, ch, nil
}
