package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sporeframework/spore/protocol"
	"google.golang.org/grpc"
)

const (
	address = "localhost:12345"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSporeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	payload := []byte("test")
	r, err := c.Send(ctx, &pb.SendTransaction{Data: payload})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("Reply: %s", r.GetMessage())
}
