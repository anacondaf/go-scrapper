package main

import (
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/internal"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/kainguyen/go-scrapper/src/core/application/http"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker/rabbitmq"
	"github.com/kainguyen/go-scrapper/src/infrastructure/serviceProvider"
	"log"
)

func main() {
	var err error

	serviceProvider.ContainerRegister()

	cfg := di.GetInstance("config").(*config.Config)

	err = internal.SetupSentry(cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	mq := di.GetInstance("rabbitmq").(*rabbitmq.RabbitMq)

	// Declare a queue to send message to
	_, err = mq.DeclareQueue(rabbitmq.QueueObject{
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

	server, err := http.NewHttpServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
