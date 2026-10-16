/*
ЗАДАЧА 13. ДОСТАВЛЯТЬ СООБЩЕНИЯ ВСЕМ ПОЛУЧАТЕЛЯМ

Что необходимо доработать:
  - Если канал получателя заполнен, убрать самое старое сообщение и положить новое.
  - Сейчас сигналы отключения получателей не обрабатываются. Нужно читать ID
    из disconnected и больше не отправлять сообщения этим получателям.
*/
package tasks

import "context"

type Message struct{}

type Recipient struct {
	ID       string
	Messages chan Message
}

func Broadcast(ctx context.Context, input <-chan Message, recipients []Recipient, disconnected <-chan string) {
	for message := range input {
		for _, recipient := range recipients {
			recipient.Messages <- message
		}
	}
}
