package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/thiendsu2303/go-grpc/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		certFile := "./ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Fail to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5*time.Second)
}
