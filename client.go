package main

import (
	"log"
	"context"
	
	"google.golang.org/grpc"

	"github.com/sithumonline/rpc-one/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("coud not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client",
	}

	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", res.Body)

}
