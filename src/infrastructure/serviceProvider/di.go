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
	"log"
	"reflect"
)

func RegisterGrpcServiceServer() {
	_, _ = di.RegisterBean("postServiceServer", reflect.TypeOf((*grpcservice.PostServiceServer)(nil)))
}

func ContainerRegister() {
	_, _ = di.RegisterBean("postHandler", reflect.TypeOf((*post.PostHandler)(nil)))
	_, _ = di.RegisterBean("postService", reflect.TypeOf((*service.PostService)(nil)))

	config, err := config.LoadConfig(utils.GetWorkDirectory())
	if err != nil {
		log.Fatalf("Error When Loading Configuration: %v\n", err)
	}

	_, err = di.RegisterBeanInstance("config", config)
	if err != nil {
		panic(err)
	}

	_, err = di.RegisterBeanFactory("webScraper", di.Singleton, func(context.Context) (interface{}, error) {
		scraper := webScraping.NewWebScraper(config, colly.AllowedDomains("vnexpress.net"))
		return scraper, nil
	})
	if err != nil {
		panic(err)
	}

	_, err = di.RegisterBeanFactory("db", di.Singleton, func(context.Context) (interface{}, error) {
		var newDB persistence.IDBConn = db.NewPostgresDB(config)

		db, err := newDB.DBConn()
		if err != nil {
			return nil, err
		}

		return db, nil
	})
	if err != nil {
		panic(err)
	}

	var redisService persistence.IRedisService = cache.NewRedisCacheService(config)
	if err != nil {
		log.Fatalf("Error When Loading Configuration: %v\n", err)
	}

	_, err = di.RegisterBeanInstance("redis", redisService)
	if err != nil {
		panic(err)
	}

	// Register rabbitmq producer & consumer
	mq, err := rabbitmq.NewRabbitMq(config)
	if err != nil {
		log.Fatalf("Error When Connect To RabbitMQ: %v\n", err)
	}

	_, err = di.RegisterBeanInstance("rabbitmq", mq)
	if err != nil {
		panic(err)
	}

	_, err = di.RegisterBeanFactory("producer", di.Singleton, func(ctx context.Context) (interface{}, error) {
		var producer = rabbitmq.NewProducer(mq)
		return producer, nil
	})
	if err != nil {
		panic(err)
	}

	_, err = di.RegisterBeanFactory("consumer", di.Singleton, func(ctx context.Context) (interface{}, error) {
		var consumer = rabbitmq.NewConsumer(mq)
		return consumer, nil
	})
	if err != nil {
		panic(err)
	}

	_, err = di.RegisterBeanFactory("websocket", di.Singleton, func(ctx context.Context) (interface{}, error) {
		var websocket = wss.NewWebsocket(redisService)
		return websocket, nil
	})
	if err != nil {
		panic(err)
	}

	RegisterGrpcServiceServer()

	_ = di.InitializeContainer()
}
