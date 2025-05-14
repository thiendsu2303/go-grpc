package main

import (
	"context"
	"log"

	pb "github.com/thiendsu2303/go-grpc/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Could not avg: %v", err)
	}
	for _, req := range []*pb.AvgRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
		{
			Number: 5,
		},
	} {
		log.Printf("Sending %v\n", req)
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not avg: %v", err)
	}
	log.Printf("Response: %v", res.Result)
}
