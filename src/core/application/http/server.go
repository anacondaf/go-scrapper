package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post/service"
	"github.com/kainguyen/go-scrapper/src/core/application/http/route"
	"reflect"
)

func init() {
	_, _ = di.RegisterBean("postHandler", reflect.TypeOf((*post.PostHandler)(nil)))
	_, _ = di.RegisterBean("postService", reflect.TypeOf((*service.PostService)(nil)))
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

	// Default config
	app.Use(cors.New())

	v1 := app.Group("/api/v1")

	route.Route(v1)

	s.app = app
}

func (s *HttpServer) StartApp(address string) error {
	return s.app.Listen(address)
}
