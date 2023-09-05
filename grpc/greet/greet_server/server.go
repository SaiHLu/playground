package main

import (
	"context"
	"fmt"
	"grpc/greet/greetpb"
	"io"
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

func (server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	var result string
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatal(err)
		}

		result += fmt.Sprintf("Hello %s\n", req.Greeting.FirstName)
	}
}

func (server) GreetEveryOne(stream greetpb.GreetService_GreetEveryOneServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal(err)
		}

		firstName := req.Greeting.FirstName
		result := "Hello " + firstName

		if err = stream.Send(&greetpb.GreetEveryOneResponse{
			Result: result,
		}); err != nil {
			log.Fatal(err)
		}
	}

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
