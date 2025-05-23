package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/thiendsu2303/go-grpc/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Fail to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	doSqrt(c, -2)
}
