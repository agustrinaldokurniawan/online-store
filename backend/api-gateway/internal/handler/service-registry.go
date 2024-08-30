package handler

import (
	"context"
	"log"

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
		var err error
		switch svc.Name {
		case "UserService":
			err = pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, svc.Address, opts)
		case "ProductService":
			err = pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, svc.Address, opts)
		case "OrderService":
			err = pb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, svc.Address, opts)
		default:
			log.Printf("No handler implemented for service %s", svc.Name)
			continue
		}
		if err != nil {
			return err
		}
		log.Printf("Registered service %s at %s", svc.Name, svc.Address)
	}
	return nil
}
