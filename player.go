package main

import (
	"context"
	"log"
	"time"

	"github.com/kanataxa/grpc-example/server"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	log.Print("Start Player")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := server.NewRelayMediatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetInput(ctx, &server.GetInputRequest{Value: 10})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("OK Empty Response: [%v]", r)
}
