package rabbitmq

import (
	"errors"
	"github.com/kainguyen/go-scrapper/src/utils/constants"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	rabbitmq *RabbitMq
}

func NewConsumer(rabbitmq *RabbitMq) *Consumer {
	return &Consumer{
		rabbitmq,
	}
}

func (c *Consumer) Consume(routingKey string, consumer string) (<-chan amqp.Delivery, error) {
	if routingKey == constants.EMPTY_STRING {
		return nil, errors.New("[messageBroker.Publish]: routingKey is required")
	}

	return c.rabbitmq.Channel.Consume(
		routingKey, // queue
		consumer,   // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
}
