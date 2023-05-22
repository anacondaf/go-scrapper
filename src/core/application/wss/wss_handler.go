package wss

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"log"
)

type Websocket struct {
	Room         *Room
	cacheService persistence.IRedisService
}

func NewWebsocket(cacheService persistence.IRedisService) *Websocket {
	room := NewRoom(uuid.New(), cacheService)

	go room.Run()
	fmt.Println("Establish websocket")

	return &Websocket{
		Room:         room,
		cacheService: cacheService,
	}
}

func (w *Websocket) upgradeWebsocket() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			log.Println("socket is connected")

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
