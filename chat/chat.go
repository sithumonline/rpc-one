package chat

import(
	"log"
	"context"
)


type Server struct {}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message bod form clinet : %s", message.Body)
	return &Message{Body: "Hello from server"}, nil
}
