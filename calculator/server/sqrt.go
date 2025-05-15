package main

import (
	"context"
	"log"
	"math"

	pb "github.com/thiendsu2303/go-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function was invoked with %v\n", in)

	if in.Number < 0 {
		return nil, status.Error(codes.InvalidArgument, "Number is not a positive number")
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(in.Number)),
	}, nil
}
