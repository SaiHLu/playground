package main

import (
	"context"
	"fmt"
	"grpc/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedSumServiceServer
}

func (s server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Request: %+v\n", req)
	res := &calculatorpb.SumResponse{
		SumResult: req.FirstNumber + req.SecondNumber,
	}

	return res, nil
}

func main() {
	fmt.Println("Starting the server...")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
