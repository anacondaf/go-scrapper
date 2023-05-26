package main

import (
	"flag"
	"fmt"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/config"
	grpcservice "github.com/kainguyen/go-scrapper/src/core/application/grpc/services/register"
	"github.com/kainguyen/go-scrapper/src/core/application/http"
	"github.com/kainguyen/go-scrapper/src/infrastructure/apm"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker/rabbitmq"
	"github.com/kainguyen/go-scrapper/src/infrastructure/serviceProvider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

var (
	// Flag
	grpcPort = flag.Int("p", 8000, "grpc port")
)

func main() {
	wg := sync.WaitGroup{}

	flag.Parse()

	var err error

	serviceProvider.ContainerRegister()

	cfg := di.GetInstance("config").(*config.Config)

	err = apm.SetupSentry(cfg)
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

	wg.Add(1)

	go startGrpcServer(&wg)

	wg.Add(1)

	go startHttpServer(&wg)

	wg.Wait()
}

func startHttpServer(wg *sync.WaitGroup) {
	defer wg.Done()

	server, err := http.NewHttpServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func startGrpcServer(wg *sync.WaitGroup) {
	defer wg.Done()

	hostname := "localhost:"
	address := fmt.Sprintf("%v%d", hostname, *grpcPort)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	grpcservice.RegisterServices(grpcServer)

	reflection.Register(grpcServer)

	fmt.Printf("GRPC Server is served at %s\n", address)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Cannot start grpc server: %v", err)
	}
}
