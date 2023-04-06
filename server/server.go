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
	s := grpc.NewServer() // создает новый сервер gRPC.
	pb.RegisterMyServiceServer(s, &server{})
	//  регистрирует наш серверный объект, реализующий интерфейс,
	// определенный в нашем protobuf-файле, в нашем сервере.
	// В данном случае мы регистрируем серверный объект типа
	// server, который реализует метод SayHello определенный в
	// нашем protobuf-определении. Это позволяет нашему серверу
	// обрабатывать входящие запросы от клиентов через методы,
	// определенные в нашем protobuf-определении.
	log.Printf("Server started on port %v", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
