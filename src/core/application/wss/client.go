package wss

import (
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"log"
)

type Client struct {
	Id      uuid.UUID `json:"id,omitempty"`
	socket  *websocket.Conn
	Receive chan []byte `json:"-"`
	Room    *Room       `json:"-"`
}

func NewClient(socket *websocket.Conn, room *Room) *Client {
	id, _ := uuid.NewUUID()

	return &Client{
		Id:      id,
		socket:  socket,
		Receive: make(chan []byte),
		Room:    room,
	}
}

func (c *Client) read() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		log.Printf("new message: %s", msg)
	}
}

func (c *Client) write() {
	defer c.socket.Close()

	for msg := range c.Receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
