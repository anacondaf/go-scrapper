package serviceProvider

import (
	"context"
	"github.com/gocolly/colly"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post/service"
	"github.com/kainguyen/go-scrapper/src/infrastructure/persistence/cache"
	"github.com/kainguyen/go-scrapper/src/infrastructure/persistence/db"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"github.com/kainguyen/go-scrapper/src/utils"
	"log"
	"reflect"
)

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

	_, err = di.RegisterBeanFactory("cache", di.Singleton, func(ctx context.Context) (interface{}, error) {
		var newCache persistence.ICacheService = cache.NewRedisCacheService(config)
		return newCache, nil
	})
	if err != nil {
		panic(err)
	}

	_ = di.InitializeContainer()
}
