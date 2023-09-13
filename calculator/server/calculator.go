package main

import (
	"context"
	"github.com/shyamchandranmec/go-grpc/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	log.Println("Sum was invoked with", req.Num1, req.Num2)
	res := &proto.CalculatorResponse{Result: req.Num1 + req.Num2}
	return res, nil
}
