package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		log.Printf("Received %v\n", req)
		res += fmt.Sprintf("Hello %s, ", req.FirstName)
	}
}
