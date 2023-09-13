package main

import (
	"context"
	"github.com/shyamchandranmec/go-grpc/calculator/proto"
	"log"
)

func sum(c proto.CalculatorServiceClient) {
	log.Println("Sum was invoked")
	res, err := c.Sum(context.Background(), &proto.CalculatorRequest{
		Num1: 3,
		Num2: 10,
	})

	if err != nil {
		log.Fatalln("Could not sum ", err)
	}
	log.Println("Sum: ", res.Result)
}
