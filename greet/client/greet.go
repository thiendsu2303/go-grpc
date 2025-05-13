package main

import (
	"log"

	"context"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Thien",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Response: %v", res)
}
