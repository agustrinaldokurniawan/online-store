package main

import (
	"log"
	"net"
	"product-service/internal/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the server with the gRPC server
	handler.RegisterServices(s)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Listen on port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	// Start the gRPC server
	log.Println("Starting gRPC server on :50052")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
