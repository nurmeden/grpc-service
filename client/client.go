package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "grpc-service/myservice"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	name := "Alice"
	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Printf("Response: %s", resp.Message)
}
