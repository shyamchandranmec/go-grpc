package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/shyamchandranmec/go-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Long greet invoked")
	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Request stream closed")
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatal("Error in recieving stream ", err)

		}
		res += fmt.Sprintf("Hello %s! \n", req.FirstName)
	}
}
