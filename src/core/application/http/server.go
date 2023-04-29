package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/goioc/di"
	_ "github.com/kainguyen/go-scrapper/docs"
	"github.com/kainguyen/go-scrapper/src/core/application/http/route"
	"github.com/kainguyen/go-scrapper/src/core/application/wss"
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
	app       *fiber.App
	websocket *wss.Websocket `di.inject:"websocket"`
}

func NewHttpServer() (*HttpServer, error) {
	websocket := di.GetInstance("websocket").(*wss.Websocket)

	server := &HttpServer{
		websocket: websocket,
	}

	server.setupApp()

	return server, nil
}

func (s *HttpServer) setupApp() {
	app := fiber.New(fiber.Config{AppName: "go-scrapper"})

	// Default config
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/ws", s.websocket.UpgradeWebsocket(), s.websocket.Handler())

	v1 := app.Group("/api/v1")

	route.Route(v1)

	s.app = app
}

func (s *HttpServer) StartApp(address string) error {
	return s.app.Listen(address)
}
