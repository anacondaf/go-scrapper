package rabbitmq

import (
	"github.com/kainguyen/go-scrapper/src/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type RabbitMq struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	logger     *zerolog.Logger
}

func NewRabbitMq(config *config.Config, logger *zerolog.Logger) (*RabbitMq, error) {
	var connectionString = config.Rabbitmq.ConnectionString

	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	logger.Info().Msg("Connect RabbitMQ Success")

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
