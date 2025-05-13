package main

import (
	"fmt"
	"log"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)
	for i := 0; i < 10; i++ {
		stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s, number: %d", in.FirstName, i),
		})
	}
	return nil
}
