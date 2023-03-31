package rabbitmq

import "encoding/json"

type MessageQueue struct {
	Body    string `json:"body"`
	Pattern string `json:"pattern"`
	Age     string `json:"age"`
	Data    string `json:"data"`
}

func NewMessageQueue(body, pattern, age, data string) *MessageQueue {
	return &MessageQueue{
		body, pattern, age, data,
	}
}

func (m *MessageQueue) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(m)

	if err != nil {
		return nil, err
	}
	return bytes, err
}
