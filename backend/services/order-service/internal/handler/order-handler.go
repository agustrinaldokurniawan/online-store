package handler

import (
	"context"
	pb "order-service/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func NewOrderServer() pb.OrderServiceServer {
	return &server{}
}

func (s *server) GetMyOrder(ctx context.Context, req *pb.GetMyOrderRequest) (*pb.GetMyOrderResponse, error) {
	// Implement your logic to get order information
	response := &pb.GetMyOrderResponse{
		UserId:    "user_id_1",
		OrderId:   "order_id_1",
		ProductId: "product_id_1",
		Quantity:  1,
		Price:     10.99,
		Status:    "pending",
	}
	return response, nil
}

func RegisterServices(s *grpc.Server) {
	pb.RegisterOrderServiceServer(s, NewOrderServer())
}
