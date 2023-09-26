package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var PORT = ":8080"

// StartgGRPC starts the gRPC server
func StartgGRPCServer() {
	fmt.Println("Starting gRPC server...")
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("could not listen to port %s: %v", PORT, err)
	}
	grpcServer := grpc.NewServer()
	RegisterPersonServer(grpcServer, &personServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("could not start gRPC server: %v", err)
	}
	fmt.Println("gRPC server started...")
}
