package serviceProvider

import (
	"context"
	"github.com/gocolly/colly"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	grpcservice "github.com/kainguyen/go-scrapper/src/core/application/grpc/services"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post/service"
	"github.com/kainguyen/go-scrapper/src/core/application/wss"
	"github.com/kainguyen/go-scrapper/src/infrastructure/messageBroker/rabbitmq"
	"github.com/kainguyen/go-scrapper/src/infrastructure/persistence/cache"
	"github.com/kainguyen/go-scrapper/src/infrastructure/persistence/db"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"github.com/kainguyen/go-scrapper/src/utils"
	"github.com/rs/zerolog"
	"reflect"
)

func RegisterGrpcServiceServer() {
	_, _ = di.RegisterBean("postServiceServer", reflect.TypeOf((*grpcservice.PostServiceServer)(nil)))
}

func RegisterLogger(logger *zerolog.Logger) error {
	_, err := di.RegisterBeanInstance("logger", logger)
	if err != nil {
		return err
	}

	return nil
}

func ContainerRegister(logger *zerolog.Logger) error {
	_, _ = di.RegisterBean("postHandler", reflect.TypeOf((*post.PostHandler)(nil)))
	_, _ = di.RegisterBean("postService", reflect.TypeOf((*service.PostService)(nil)))

	config, err := config.LoadConfig(utils.GetWorkDirectory(), logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error When Loading Configuration")
	}

	_, err = di.RegisterBeanInstance("config", config)
	if err != nil {
		return err
	}

	_, err = di.RegisterBeanFactory("webScraper", di.Singleton, func(context.Context) (interface{}, error) {
		scraper := webScraping.NewWebScraper(config, logger, colly.AllowedDomains("vnexpress.net"))
		return scraper, nil
	})
	if err != nil {
		return err
	}

	_, err = di.RegisterBeanFactory("db", di.Singleton, func(context.Context) (interface{}, error) {
		var newDB persistence.IDBConn = db.NewPostgresDB(config, logger)

		db, err := newDB.DBConn()
		if err != nil {
			return nil, err
		}

		return db, nil
	})
	if err != nil {
		return err
	}

	var redisService persistence.IRedisService = cache.NewRedisCacheService(config, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error When Loading Configuration")
	}

	_, err = di.RegisterBeanInstance("redis", redisService)
	if err != nil {
		return err
	}

	// Register rabbitmq producer & consumer
	mq, err := rabbitmq.NewRabbitMq(config, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error When Connect To RabbitMQ")
	}

	_, err = di.RegisterBeanInstance("rabbitmq", mq)
	if err != nil {
		return err
	}

	_, err = di.RegisterBeanFactory("producer", di.Singleton, func(ctx context.Context) (interface{}, error) {
		var producer = rabbitmq.NewProducer(mq)
		return producer, nil
	})
	if err != nil {
		return err
	}

	_, err = di.RegisterBeanFactory("consumer", di.Singleton, func(ctx context.Context) (interface{}, error) {
		var consumer = rabbitmq.NewConsumer(mq)
		return consumer, nil
	})
	if err != nil {
		return err
	}

	_, err = di.RegisterBeanFactory("websocket", di.Singleton, func(ctx context.Context) (interface{}, error) {
		var websocket = wss.NewWebsocket(redisService, logger)
		return websocket, nil
	})
	if err != nil {
		return err
	}

	RegisterGrpcServiceServer()

	err = RegisterLogger(logger)
	if err != nil {
		return err
	}

	_ = di.InitializeContainer()

	return nil
}
