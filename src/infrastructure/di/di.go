package di

import (
	"context"
	"github.com/gocolly/colly"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/infrastructure/persistence/db"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"github.com/kainguyen/go-scrapper/src/utils"
	"log"
)

func init() {
	config, err := config.LoadConfig(utils.GetWorkDirectory())
	if err != nil {
		log.Fatalf("Error When Loading Configuration: %v\n", err)
	}

	_, err = di.RegisterBeanInstance("config", config)
	if err != nil {
		panic(err)
	}

	_, err = di.RegisterBeanFactory("webScraper", di.Singleton, func(context.Context) (interface{}, error) {
		scraper := webScraping.NewWebScraper(colly.AllowedDomains("vnexpress.net"))
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

	_ = di.InitializeContainer()
}
