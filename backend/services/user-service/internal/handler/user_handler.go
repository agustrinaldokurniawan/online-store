package handler

import (
	"context"
	"log"

	pb "user-service/pb"

	"google.golang.org/grpc"
)

// server implements the UserServiceServer interface
type server struct {
	pb.UnimplementedUserServiceServer
}

// NewServer creates a new instance of server
func NewServer() pb.UserServiceServer {
	return &server{}
}

// Implement the methods defined in your proto file here
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Log the request details
	log.Printf("Received GetUser request with ID: %s", req.GetId())

	// Implement your logic to get user information
	response := &pb.GetUserResponse{
		Id:    req.GetId(),
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Log the response details
	log.Printf("Returning GetUser response: %+v", response)

	return response, nil
}

// Implement GetUserProfile method
func (s *server) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	// Log the request details
	log.Printf("Received GetUserProfile request with ID: %s", req.GetId())

	response := &pb.GetUserProfileResponse{
		Id:      req.GetId(),
		Name:    "John Doe",
		Email:   "john.doe@example.com",
		Phone:   "1234567890",
		Address: "123 Main St",
	}

	// Log the response details
	log.Printf("Returning GetUserProfile response: %+v", response)

	return response, nil
}

// RegisterServices registers the gRPC services with the server
func RegisterServices(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, NewServer())
}
