package main

import (
	"github.com/shyamchandranmec/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Unable to dial ", addr, err)
	}
	log.Println("Dial successful")
	defer conn.Close()
	c := proto.NewCalculatorServiceClient(conn)
	sum(c)
}
