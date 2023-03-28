package main

import (
	"context"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/kainguyen/go-scrapper/src/core/application/http"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker"
	"github.com/kainguyen/go-scrapper/src/infrastructure/serviceProvider"
	"log"
)

func main() {
	serviceProvider.ContainerRegister()

	config := di.GetInstance("config").(*config.Config)

	rabbitmq, err := messageBroker.NewRabbitMq(config)
	if err != nil {
		log.Fatalf("%v", err)
	}

	producer := messageBroker.NewProducer(rabbitmq)

	// Declare a queue to send message to
	queue, err := producer.DeclareQueue(messageBroker.QueueObject{
		QueueName:  "hello",
		Durable:    false,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	message := "Hello World"

	// Publish message to queue q with name "hello"
	err = producer.Publish(context.Background(), queue.Name, message)
	if err != nil {
		log.Fatalf("%v", err)
	}

	server, err := http.NewHttpServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
