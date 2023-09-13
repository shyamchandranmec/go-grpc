package main

import (
	"github.com/shyamchandranmec/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen to addrr ", addr, err)
	}
	log.Println("Listening to address ", addr)
	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &Server{})
	if s.Serve(lis); err != nil {
		log.Fatalln("Unable to server ", err)
	}
}
