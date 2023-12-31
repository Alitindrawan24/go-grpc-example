package main

import (
	"context"
	"grpc/grpc/helloworld/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedHelloWorldServer
}

// SayHello implements a server
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s *server) Factorial(ctx context.Context, in *pb.FactorialRequest) (*pb.FactorialResponse, error) {
	log.Printf("Received: %v", in.GetNumber())
	fact := int64(1)
	for i := int64(1); i <= in.GetNumber(); i++ {
		fact = fact * i
	}

	return &pb.FactorialResponse{Number: int64(fact)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
