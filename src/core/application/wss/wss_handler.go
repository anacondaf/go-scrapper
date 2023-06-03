package wss

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/rs/zerolog"
)

type Websocket struct {
	Room         *Room
	cacheService persistence.IRedisService
	logger       *zerolog.Logger
}

func NewWebsocket(cacheService persistence.IRedisService, logger *zerolog.Logger) *Websocket {
	room := NewRoom(uuid.New(), cacheService)

	go room.Run()

	logger.Info().Msg("Establish websocket")

	return &Websocket{
		Room:         room,
		cacheService: cacheService,
		logger:       logger,
	}
}

func (w *Websocket) upgradeWebsocket() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			w.logger.Info().Msg("Socket Is Connected")

			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	}
}

func (w *Websocket) JoinRoom() fiber.Handler {
	w.upgradeWebsocket()

	return websocket.New(func(socket *websocket.Conn) {
		client := NewClient(socket, w.Room)

		// Join room
		w.Room.Join <- client

		go client.write()
		client.read()
	})
}

func (w *Websocket) ListAllClients() fiber.Handler {
	w.upgradeWebsocket()

	return websocket.New(func(socket *websocket.Conn) {
		w.Room.ListClients()
	})
}
