package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "api-gateway/pb"
)

// ServiceConfig holds the configuration for each service.
type ServiceConfig struct {
	Name    string
	Address string
}

// RegisterServices registers all services with the mux based on the given service configurations.
func RegisterServices(ctx context.Context, mux *runtime.ServeMux, services []ServiceConfig) error {
	for _, svc := range services {
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, svc.Address, opts)
		if err != nil {
			return err
		}
		log.Printf("Registered service %s at %s", svc.Name, svc.Address)
	}
	return nil
}

func main() {
	// Define service configurations
	services := []ServiceConfig{
		{Name: "UserService", Address: "user-service:50051"},
		// Add more services here
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Register services
	if err := RegisterServices(ctx, mux, services); err != nil {
		log.Fatalf("Failed to register services: %v", err)
	}

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
