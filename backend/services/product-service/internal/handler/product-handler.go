package handler

import (
	"context"
	pb "product-service/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func NewProfileServer() pb.ProductServiceServer {
	return &server{}
}

func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	// Implement your logic to get product information
	response := &pb.GetProductResponse{
		Id:          req.GetId(),
		Name:        "Product Name",
		Price:       10.99,
		Description: "Product Description",
	}
	return response, nil
}

func RegisterServices(s *grpc.Server) {
	pb.RegisterProductServiceServer(s, NewProfileServer())
}
