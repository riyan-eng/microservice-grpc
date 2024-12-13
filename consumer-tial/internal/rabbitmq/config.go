package rabbitmq

const (
	exchangeName     = "delayed_exchange"
	exchangeType     = "x-delayed-message"
	queueName        = "task_queue"
	routingKey       = "notifications"
	initialDelayMs   = 5000
	retryDelayFactor = 2
	maxRetries       = 5
)
