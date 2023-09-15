package main

import (
	"fmt"
	"github.com/shyamchandranmec/go-grpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	log.Println("Greet many times invoked")
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s , number %d", in.FirstName, i)
		stream.Send(&proto.GreetResponse{Result: res})
	}
	return nil
}
