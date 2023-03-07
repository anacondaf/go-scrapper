package http

import (
	"context"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/core/application/http/blog"
	"github.com/kainguyen/go-scrapper/src/core/application/http/blog/service"
	"github.com/kainguyen/go-scrapper/src/infrastructure/persistence/db"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"reflect"
)

func init() {
	_, _ = di.RegisterBean("blogHandler", reflect.TypeOf((*blog.Handler)(nil)))
	_, _ = di.RegisterBean("blogService", reflect.TypeOf((*service.BlogService)(nil)))

	di.RegisterBeanFactory("webScraper", di.Singleton, func(context.Context) (interface{}, error) {
		scraper := webScraping.NewWebScraper(colly.AllowedDomains("vnexpress.net"))
		return scraper, nil
	})

	_, err := di.RegisterBeanFactory("db", di.Singleton, func(context.Context) (interface{}, error) {
		var newDB persistence.IDBConn = db.NewPostgresDB()

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

type HttpServer struct {
	app *fiber.App
}

func NewHttpServer() (*HttpServer, error) {
	server := &HttpServer{}

	server.setupApp()

	return server, nil
}

func (s *HttpServer) setupApp() {
	app := fiber.New(fiber.Config{AppName: "go-scrapper"})

	v1 := app.Group("/api/v1")

	blogRouter := v1.Group("blogs")

	var blogHandler = di.GetInstance("blogHandler").(*blog.Handler).CreatePost()
	blogRouter.Post("/", blogHandler)

	s.app = app
}

func (s *HttpServer) StartApp(address string) error {
	return s.app.Listen(address)
}
