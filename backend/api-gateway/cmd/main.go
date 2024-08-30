package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"api-gateway/internal/handler"
)

func main() {
	// Define service configurations
	services := []handler.ServiceConfig{
		{Name: "UserService", Address: "user-service:50051"},
		{Name: "ProductService", Address: "product-service:50052"},
		{Name: "OrderService", Address: "order-service:50053"},
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Register services
	if err := handler.RegisterServices(ctx, mux, services); err != nil {
		log.Fatalf("Failed to register services: %v", err)
	}

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
