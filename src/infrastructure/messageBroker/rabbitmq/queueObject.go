package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type QueueObject struct {
	QueueName  string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}
