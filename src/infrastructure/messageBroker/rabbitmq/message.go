package rabbitmq

/* Technical Note:
* To send message to queue with the message pattern (aka message name) such that NestJS service can receive, we need to construct a struct contain a pattern field
* For example:
* @EventPattern('post') // NestJS will only consume message with 'post' pattern(or name)
* Then:
* In Go service, we need to construct the message with the pattern field as below:
* Body: Message{Pattern: "post", Data: <Our message's data>}
* | REFERENCES: https://stackoverflow.com/questions/70513069/how-to-pass-a-message-from-go-and-consuming-it-from-nestjs-with-rabbitmq/70514684 |
 */

type Message struct {
	Pattern string      `json:"pattern"`
	Data    interface{} `json:"data"`
}

func NewMessage(pattern string, data interface{}) *Message {
	return &Message{
		pattern, data,
	}
}
