package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/kainguyen/go-scrapper/docs"
	"github.com/kainguyen/go-scrapper/src/core/application/http/route"
	"github.com/rs/zerolog"
)

//	@title			Fiber Example API
//	@version		1.0
//	@description	This is a sample swagger for Fiber
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:3000
//	@BasePath		/api/v1

type HttpServer struct {
	logger *zerolog.Logger
	app    *fiber.App
}

func NewHttpServer(logger *zerolog.Logger) (*HttpServer, error) {

	server := &HttpServer{
		logger: logger,
	}

	server.setupApp()

	return server, nil
}

func (s *HttpServer) setupApp() {
	app := fiber.New(fiber.Config{AppName: "go-scrapper"})

	// Default config
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	v1 := app.Group("/api/v1")

	route.HttpRoute(v1, s.logger)

	wssRouter := app.Group("/ws")

	route.WebsocketRoute(wssRouter)

	s.app = app
}

func (s *HttpServer) StartApp(address string) error {
	return s.app.Listen(address)
}
