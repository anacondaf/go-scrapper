package main

import (
	"flag"
	"fmt"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/config"
	grpcservice "github.com/kainguyen/go-scrapper/src/core/application/grpc/services/register"
	"github.com/kainguyen/go-scrapper/src/core/application/http"
	"github.com/kainguyen/go-scrapper/src/infrastructure/apm"
	logservice "github.com/kainguyen/go-scrapper/src/infrastructure/log"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker/rabbitmq"
	"github.com/kainguyen/go-scrapper/src/infrastructure/serviceProvider"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
)

var (
	// Flag
	grpcPort = flag.Int("p", 8000, "grpc port")
	logger   = logservice.NewLogger()
)

func init() {
	err := serviceProvider.ContainerRegister(logger)
	if err != nil {
		logger.Fatal().Err(err)
	}
}

func main() {
	wg := sync.WaitGroup{}

	flag.Parse()

	var err error

	cfg := di.GetInstance("config").(*config.Config)

	err = apm.SetupSentry(cfg)
	if err != nil {
		logger.Error().Err(err)
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
		logger.Error().Err(err)
	}

	wg.Add(1)

	go startGrpcServer(&wg)

	wg.Add(1)

	go startHttpServer(&wg, logger)

	wg.Wait()
}

func startHttpServer(wg *sync.WaitGroup, logger *zerolog.Logger) {
	defer wg.Done()

	server, err := http.NewHttpServer(logger)
	if err != nil {
		logger.Error().Err(err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		logger.Error().Err(err)
	}
}

func startGrpcServer(wg *sync.WaitGroup) {
	defer wg.Done()

	hostname := "localhost:"
	address := fmt.Sprintf("%v%d", hostname, *grpcPort)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error().Err(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	grpcservice.RegisterServices(grpcServer)

	reflection.Register(grpcServer)

	logger.Info().Msg(fmt.Sprintf("GRPC Server is served at %s\n", address))

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error().Stack().Err(err).Msg(fmt.Sprintf("Cannot start grpc server"))
	}
}
