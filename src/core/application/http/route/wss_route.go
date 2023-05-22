package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/core/application/wss"
)

func WebsocketRoute(router fiber.Router) {
	wssHandler := di.GetInstance("websocket").(*wss.Websocket)

	router.Get("/room/join", wssHandler.JoinRoom())
	router.Get("/clients", wssHandler.ListAllClients())
}
