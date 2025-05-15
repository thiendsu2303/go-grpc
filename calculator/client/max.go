package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/thiendsu2303/go-grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Could not max: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}

		for _, number := range numbers {
			log.Printf("Sending %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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
				log.Fatalf("Could not receive max: %v", err)
			}
			log.Printf("Max: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
