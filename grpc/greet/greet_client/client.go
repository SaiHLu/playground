package main

import (
	"context"
	"fmt"
	"grpc/greet/greetpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("From Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	UnaryRequest(client)

	ServerStreamingRequest(client)
}

func UnaryRequest(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sai Hlaing",
			LastName:  " Lu",
		},
	}

	response, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response: ", response.Result)
}

func ServerStreamingRequest(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sai Hlaing ",
			LastName:  "Lu",
		},
	}

	stream, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Streaming...", msg.Result)
	}
}
