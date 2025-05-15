package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		log.Printf("Received %v\n", req)
		stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s, ", req.FirstName),
		})
		if err != nil {
			log.Fatalf("Error while sending response to client: %v", err)
		}
	}
}
