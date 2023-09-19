package grpc

import context "context"

type personServer struct {
	UnimplementedPersonServer
}

func (s *personServer) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello " + in.GetName()}, nil
}
