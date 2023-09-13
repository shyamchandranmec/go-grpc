package main

import (
	"context"
	"github.com/shyamchandranmec/go-grpc/greet/proto"
	"log"
)

func doGreet(c proto.GreetServiceClient) {
	log.Println("Do greet was invoked")
	res, err := c.Greet(context.Background(), &proto.GreetRequest{FirstName: "Shyam"})
	if err != nil {
		log.Fatalln("Could not greet ", err)
	}
	log.Println("Greeting: ", res.Result)
}
