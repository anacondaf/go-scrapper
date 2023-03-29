package rabbitmq

import (
	"fmt"
	"github.com/kainguyen/go-scrapper/src/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMq(config *config.Config) (*RabbitMq, error) {
	var connectionString = config.Rabbitmq.ConnectionString

	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connect RabbitMQ Success!")

	return &RabbitMq{
		Connection: conn,
		Channel:    channel,
	}, nil
}

func (mq *RabbitMq) DeclareQueue(queueObj QueueObject) (amqp.Queue, error) {
	return mq.Channel.QueueDeclare(
		queueObj.QueueName,
		queueObj.Durable,
		queueObj.AutoDelete,
		queueObj.Exclusive,
		queueObj.NoWait,
		queueObj.Args,
	)
}

// Close ...
func (mq *RabbitMq) Close() {
	mq.Connection.Close()
}
