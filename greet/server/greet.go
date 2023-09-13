package main

import (
	"context"
	"github.com/shyamchandranmec/go-grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Println("Greet was invoked  with ", in)
	return &proto.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
