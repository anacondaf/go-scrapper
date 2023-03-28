package messageBroker

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

// Close ...
func (r *RabbitMq) Close() {
	r.Connection.Close()
}
