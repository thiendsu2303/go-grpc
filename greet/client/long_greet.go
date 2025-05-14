package main

import (
	"context"
	"log"
	"time"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet function was invoked")

	reqs := []*pb.GreetRequest{
		{
			FirstName: "Thien",
		},
		{
			FirstName: "Vinh",
		},
		{
			FirstName: "Trung",
		},
		{
			FirstName: "Hien",
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Could not greet many times: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not greet many times: %v", err)
	}
	log.Printf("Response: %v", res.Result)
}
