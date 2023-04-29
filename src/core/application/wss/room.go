package wss

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/core/domain/enums"
	"log"
)

type Room struct {
	Id      uuid.UUID             `json:"id,omitempty"`
	Clients map[uuid.UUID]*Client `json:"clients,omitempty"`
	Forward chan []byte           `json:"forward,omitempty"`
	Join    chan *Client          `json:"join,omitempty"`

	redisService persistence.IRedisService
}

func NewRoom(id uuid.UUID, redisService persistence.IRedisService) *Room {
	return &Room{
		Id:           id,
		Forward:      make(chan []byte),
		Join:         make(chan *Client),
		Clients:      make(map[uuid.UUID]*Client),
		redisService: redisService,
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Join:
			r.Clients[client.Id] = client

			clientIds, err := r.GetAllClientId()
			if err != nil {
				log.Println(err)
			}

			r.redisService.Set(context.Background(), enums.WSS_CLIENTS, clientIds, 0)
		}
	}
}

func (r *Room) ListClients() {
	for i, _ := range r.Clients {
		fmt.Printf("%s\n", i)
	}
}

func (r *Room) GetAllClientId() ([]string, error) {
	var clientIds []string

	for key, _ := range r.Clients {
		clientIds = append(clientIds, key.String())
	}

	return clientIds, nil
}
