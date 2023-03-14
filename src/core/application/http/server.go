package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post"
	"github.com/kainguyen/go-scrapper/src/core/application/http/post/service"
	"reflect"
)

func init() {
	_, _ = di.RegisterBean("postHandler", reflect.TypeOf((*post.Handler)(nil)))
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

	v1 := app.Group("/api/v1")

	postRouter := v1.Group("posts")

	var postHandler = di.GetInstance("postHandler").(*post.Handler)
	postRouter.Post("/", postHandler.CreatePost())
	postRouter.Get("/", postHandler.GetPosts())

	s.app = app
}

func (s *HttpServer) StartApp(address string) error {
	return s.app.Listen(address)
}
