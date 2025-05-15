package main

import (
	"io"
	"log"

	pb "github.com/thiendsu2303/go-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")
	var maximum int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})
			if err != nil {
				log.Fatalf("Error while sending response to client: %v", err)
			}
		}
	}
}
