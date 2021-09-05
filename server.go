package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sithumonline/rpc-one/chat"
)

func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Server started on port 8080")

}

//	 protoc --go_out=plugins=grpc:chat chat.proto
