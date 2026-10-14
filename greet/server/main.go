package main

import (
	"log"
	"net"

	pb "github.com/greet-go-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	s := grpc.NewServer()

	// Register the GreetService server implementation with the gRPC server
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
