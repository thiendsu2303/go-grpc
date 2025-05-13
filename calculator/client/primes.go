package main

import (
	"context"
	"io"
	"log"

	pb "github.com/thiendsu2303/go-grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimesRequest{
		Number: 1239039284,
	}
	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not primes: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive primes: %v", err)
		}
		log.Printf("Response: %v", res.Result)
	}
}
