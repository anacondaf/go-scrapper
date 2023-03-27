package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kainguyen/go-scrapper/src/core/application/http"
	"github.com/kainguyen/go-scrapper/src/infrastructure/di"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	di.ContainerRegister()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer ch.Close()

	fmt.Println("Connect RabbitMQ Success!")

	// Declare a queue to send message to

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	// Publish message to queue q with name "hello"
	body := "Hello World!"
	err = ch.PublishWithContext(context.Background(),
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
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
