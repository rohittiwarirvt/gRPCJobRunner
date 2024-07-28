package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "rohittiwarirvt/grpc/helloworld"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworl.GreaterServer

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello again " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer()

	pb.RegisterGreeterServer(gRPCServer, &server{})
	log.Printf("Server listening at %v", list.Addr())
	if err := gRPCServer.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
