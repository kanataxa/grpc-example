package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kanataxa/grpc-example/server"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type RelayMediatorServer struct {
	server.RelayMediatorServer
}

func (s *RelayMediatorServer) GetInput(ctx context.Context, in *server.GetInputRequest) (*empty.Empty, error) {
	log.Printf("Received: %v", in.GetValue())
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server.RegisterRelayMediatorServer(s, &RelayMediatorServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
