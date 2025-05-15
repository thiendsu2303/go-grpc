package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("GreetWithDeadline function was invoked")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, status.Error(codes.Canceled, "Deadline Exceeded")
		}
	}

	time.Sleep(1 * time.Second)
	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", in.FirstName),
	}, nil
}
