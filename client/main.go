package main

import (
	"GRPC_server/chat"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("did not connect: %s", err)
	}

	//defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	//Send a message to the server every 3 seconds
	for(true) {
	
		respose, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from client"})

		if err != nil {
			log.Fatal("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", respose.Body)	

		time.Sleep(3 * time.Second) 
	}
	

}