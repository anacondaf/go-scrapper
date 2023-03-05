package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kainguyen/go-scrapper/src/core/application/api/blog"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
)

type Server struct {
	app   *fiber.App
	colly *webScraping.Colly
}

func NewHttpServer(colly *webScraping.Colly) (*Server, error) {
	server := &Server{colly: colly}

	server.setupApp()

	return server, nil
}

func (s *Server) setupApp() {
	app := fiber.New(fiber.Config{AppName: "go-scrapper"})

	v1 := app.Group("/api/v1")

	blogRouter := v1.Group("blogs")
	blogRouter.Post("/", blog.CreatePost(s.colly))

	s.app = app
}

func (s *Server) StartApp(address string) error {
	return s.app.Listen(address)
}
