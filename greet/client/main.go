package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50051"

func main() {
	// by default, gRPC uses HTTP/2 and requires TLS. For the sake of this example, we will disable TLS to avoid needing certificates.
	// we by pass the grpc.WithInsecure() option to the Dial function to disable TLS.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

}
