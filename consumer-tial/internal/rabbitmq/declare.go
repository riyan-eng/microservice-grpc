package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// DeclareTopology sets up the exchange, queue, and bindings.
func DeclareTopology(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		amqp.Table{"x-delayed-type": "direct"},
	)
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return ch.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
}
