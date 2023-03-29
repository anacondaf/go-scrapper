package main

import (
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/core/application/http"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker/rabbitmq"
	"github.com/kainguyen/go-scrapper/src/infrastructure/serviceProvider"
	"log"
)

func main() {
	serviceProvider.ContainerRegister()

	mq := di.GetInstance("rabbitmq").(*rabbitmq.RabbitMq)

	//consumer := di.GetInstance("consumer").(*rabbitmq.Consumer)

	// Declare a queue to send message to
	_, err := mq.DeclareQueue(rabbitmq.QueueObject{
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

	//msgs, err := consumer.Consume(queue.Name, "")
	//if err != nil {
	//	log.Fatalf("%v", err)
	//}
	//
	//go func() {
	//	for msg := range msgs {
	//		log.Printf("Received a message: %s\n", msg.Body)
	//	}
	//}()

	server, err := http.NewHttpServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
