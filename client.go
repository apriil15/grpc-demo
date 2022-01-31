package main

import (
	"context"
	"log"

	"github.com/apriil15/grpc-demo/chat"
	"github.com/apriil15/grpc-demo/utils"
	"google.golang.org/grpc"
)

func main() {
	// create channel
	connection, err := grpc.Dial(utils.PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect")
	}
	defer connection.Close()

	// create client
	client := chat.NewChatServiceClient(connection)

	response, err := client.SayHello(context.Background(), &chat.MessageRequest{Body: "hello from the client!"})
	if err != nil {
		log.Fatal("Error when client is calling SayHello")
	}

	log.Printf("response from server: %d", response.Result)
}
