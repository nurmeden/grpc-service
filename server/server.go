package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-service/myservice"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	log.Printf("Server started on port %v", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
