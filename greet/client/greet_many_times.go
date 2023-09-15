package main

import (
	"context"
	"github.com/shyamchandranmec/go-grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c proto.GreetServiceClient) {
	log.Println("Greet many times was invoked")
	req := &proto.GreetRequest{FirstName: "Shyam"}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalln("Error in stream greet many times")
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while reading stream ", err)

		}
		log.Printf("Greet many %s", msg.Result)
	}
}
