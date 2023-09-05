package main

import (
	"context"
	"fmt"
	"grpc/greet/greetpb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Request %+v\n", req)
	firstName := req.Greeting.FirstName

	return &greetpb.GreetResponse{
		Result: "Hello " + firstName,
	}, nil
}

func (server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.Greeting.FirstName

	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello %s, numbers %d", firstName, i)
		response := &greetpb.GreetManyTimesRespose{
			Result: result,
		}

		stream.Send(response)

		time.Sleep(time.Second * 1)
	}

	return nil
}

func main() {
	fmt.Println("Hello World")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
