package main

import (
	"context"
	"fmt"
	"log"
	"net"
	t "time"

	services "github.com/jdmini/Disys2/chat"

	grpc "google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) Broadcast(ctx context.Context, in *services.Broadcast)

func (s *Server) GetTime(ctx context.Context, in *time.GetTimeRequest) (*time.GetTimeReply, error) {
	fmt.Printf("Received GetTime request\n")
	return &time.GetTimeReply{Reply: t.Now().String()}, nil
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	time.RegisterGetCurrentTimeServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
