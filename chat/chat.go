package chat

import (
	"context"
	"fmt"
	"log"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message bod form clinet : %s", message.Body)
	replyBody := fmt.Sprintf("Hello %d from server", message.Count)
	if message.Count > 200 {
		return &Message{
			Body:  replyBody,
			Count: message.Count - 100,
		}, nil
	}
	return &Message{
		Body:  replyBody,
		Count: message.Count + 10,
	}, nil
}
