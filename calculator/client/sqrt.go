package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/thiendsu2303/go-grpc/calculator/proto"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Received an error: %v", e.Message())
			log.Printf("Received an error code: %v", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("Received an invalid argument: %v", e.Message())
				return
			}
		} else {
			log.Fatalf("Could not sqrt: %v", err)
		}
	}
	log.Printf("Sqrt: %v", res.Result)
}
