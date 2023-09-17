package main

import (
	"context"
	"io"
	"log"

	pb "github.com/shyamchandranmec/go-grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes invoked ")
	req := &pb.PrimeRequest{
		Number: 1233432432,
	}
	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatal("Error in doPrime", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while reading prime steram")
		}
		log.Println("Primes ", res.Result)
	}

}
