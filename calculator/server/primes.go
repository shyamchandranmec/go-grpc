package main

import (
	"log"

	pb "github.com/shyamchandranmec/go-grpc/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Println("Primes function was invoked with ", in)
	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number = number / divisor
		} else {
			divisor++
		}
	}
	return nil
}
