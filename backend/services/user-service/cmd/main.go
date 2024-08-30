package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"user-service/internal/handler"
)

func main() {
	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the server with the gRPC server
	handler.RegisterServices(s)

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the gRPC server
	log.Println("Starting gRPC server on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
