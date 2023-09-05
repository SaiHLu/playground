package main

import (
	"context"
	"fmt"
	"grpc/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := calculatorpb.NewSumServiceClient(conn)

	res, err := client.Sum(context.Background(), &calculatorpb.SumRequest{FirstNumber: 10, SecondNumber: 20})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response: ", res.SumResult)
}
