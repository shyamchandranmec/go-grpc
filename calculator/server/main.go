package main

import (
	"github.com/shyamchandranmec/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "0.0.0.0:50052"

type Server struct {
	proto.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen to addrr ", addr, err)
	}
	log.Println("Listening to address ", addr)

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{})
	if s.Serve(lis); err != nil {
		log.Fatalln("Unable to server ", err)
	}
}
