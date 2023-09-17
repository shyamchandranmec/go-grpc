package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/shyamchandranmec/go-grpc/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("do average invoked")
	reqs := []*pb.AvgRequest{
		{
			Num: 1,
		},
		{
			Num: 2,
		},
		{
			Num: 3,
		},
		{
			Num: 4,
		},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Println("Errorn in streaming from client")
	}
	for _, req := range reqs {
		log.Println("Sending ", req.Num)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, _ := stream.CloseAndRecv()
	fmt.Println("Average is ", res.Result)

}
