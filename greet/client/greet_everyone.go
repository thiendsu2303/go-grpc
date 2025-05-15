package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v", err)
	}
	reqs := []*pb.GreetRequest{
		{
			FirstName: "Thien",
		},
		{
			FirstName: "Anh",
		},
		{
			FirstName: "Hien",
		},
		{
			FirstName: "Hoa",
		},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Sending %v\n", req)
			stream.Send(req)
			time.Sleep(100 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response from server: %v", err)
			}
			log.Printf("Response: %v\n", res)
		}
		close(waitc)
	}()

	<-waitc
}
