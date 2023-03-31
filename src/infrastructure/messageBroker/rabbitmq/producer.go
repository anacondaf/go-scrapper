package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kainguyen/go-scrapper/src/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	rabbitmq *RabbitMq
}

func NewProducer(rabbitmq *RabbitMq) *Producer {
	return &Producer{
		rabbitmq,
	}
}

func (p *Producer) Publish(context context.Context, routingKey string, message interface{}) error {
	if routingKey == utils.EMPTY_STRING {
		return errors.New("[messageBroker.Publish]: routingKey is required")
	}

	//var b bytes.Buffer
	//
	//if err := gob.NewEncoder(&b).Encode(message); err != nil {
	//	return errors.New(fmt.Sprintf("[messageBroker.Publish]: %v", err))
	//}

	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.rabbitmq.Channel.PublishWithContext(context,
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		},
	)
	if err != nil {
		return errors.New(fmt.Sprintf("[messageBroker.Publish]: %v", err))
	}

	fmt.Printf("Publish to queue %s success\n", routingKey)

	return nil
}
