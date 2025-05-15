package main

import (
	"context"
	"log"
	"time"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Thien",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Received an deadline exceeded error: %v", e.Message())
				return
			} else {
				log.Fatalf("Received an error: %v", e.Message())
			}
		} else {
			log.Fatalf("Could not greet with deadline: %v", err)
		}
	}

	log.Printf("GreetWithDeadline: %v\n", res.Result)
}
