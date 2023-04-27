package wss

import "github.com/google/uuid"

type Client struct {
	Id uuid.UUID `json:"id"`
}
