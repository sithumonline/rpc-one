package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/sithumonline/rpc-one/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("coud not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body:  "Hello from the client",
		Count: 100,
	}

	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", res.Body)

	for {
		time.Sleep(time.Second * 5)
		var msg *chat.Message
		replyBody := fmt.Sprintf("Hello %d from client", res.Count)
		if res.Count > 200 {
			msg = &chat.Message{
				Body:  replyBody,
				Count: res.Count - 100,
			}
		}
		msg = &chat.Message{
			Body:  replyBody,
			Count: res.Count + 10,
		}
		res, err = c.SayHello(context.Background(), msg)
		if err != nil {
			log.Fatalf("error when calling SayHello at for: %s", err)
		}
		log.Printf("Response from server at for: %s", res.Body)
	}

}
