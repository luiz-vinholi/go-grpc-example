package main

import (
	"context"
	"log"

	grpcserver "go-grpc-example/interfaces/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcserver.NewPersonClient(conn)
	hello, err := client.SayHello(context.TODO(), &grpcserver.HelloRequest{Name: "Thrawn"})
	if err != nil {
		log.Fatalf("could not say hello: %v", err)
	}
	log.Print(hello)
}
