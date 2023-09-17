package main

import (
	"io"
	"log"
	pb "github.com/shyamchandranmec/go-grpc/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg server invoked")
	res := 0.0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Request stream closed")
			res /= float64(count)
			return stream.SendAndClose(&pb.AvgResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatal("Error in reading stream")
		}
		count++
		log.Println("Before Res , Count", res, count)
		res += float64(req.Num)
		log.Println("After Res , Count", res, count)
	}
}
