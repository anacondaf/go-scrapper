package messageBroker

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/kainguyen/go-scrapper/src/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	rabbitmq *RabbitMq
}

type QueueObject struct {
	QueueName  string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

func NewProducer(rabbitmq *RabbitMq) *Producer {
	return &Producer{
		rabbitmq,
	}
}

func (p *Producer) DeclareQueue(queueObj QueueObject) (amqp.Queue, error) {
	return p.rabbitmq.Channel.QueueDeclare(
		queueObj.QueueName,
		queueObj.Durable,
		queueObj.AutoDelete,
		queueObj.Exclusive,
		queueObj.NoWait,
		queueObj.Args,
	)
}

func (p *Producer) Publish(context context.Context, routingKey string, message interface{}) error {
	if routingKey == utils.EMPTY_STRING {
		return errors.New("[messageBroker.Publish]: routingKey is required")
	}

	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(message); err != nil {
		return errors.New(fmt.Sprintf("[messageBroker.Publish]: %v", err))
	}

	err := p.rabbitmq.Channel.PublishWithContext(context,
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b.Bytes(),
		},
	)
	if err != nil {
		return errors.New(fmt.Sprintf("[messageBroker.Publish]: %v", err))
	}

	return nil
}
