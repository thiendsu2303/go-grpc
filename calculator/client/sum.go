package main

import (
	"context"
	"log"

	pb "github.com/thiendsu2303/go-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}
	log.Printf("Sum: %v", res.Result)
}
