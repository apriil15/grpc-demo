package main

import (
	"context"
	"log"
	"net"

	"github.com/apriil15/grpc-demo/chat"
	"github.com/apriil15/grpc-demo/utils"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", utils.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc servers
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &Server{})

	log.Printf("gRPC server is listening on%s", utils.PORT)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("gRPC server error")
	}
}

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *chat.MessageRequest) (*chat.MessageResponse, error) {
	log.Print(message.Body)

	return &chat.MessageResponse{Result: 123}, nil
}
