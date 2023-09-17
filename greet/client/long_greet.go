package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shyamchandranmec/go-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Do long greet was invoked")

	reqs := []*pb.GreetRequest{
		{
			FirstName: "Shyam",
		},
		{
			FirstName: "Sharath",
		},
		{
			FirstName: "Leela",
		},
		{
			FirstName: "Chandran",
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatal("Error in sending stream request LongGreet")
	}
	for _, req := range reqs {
		log.Println("Sending req ", req)
		stream.Send(req)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Error in receiveing long greet response")
	}
	log.Println("Long greet result \n", res.Result)
}
