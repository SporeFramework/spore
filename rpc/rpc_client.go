package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
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
	/*
		payload := []byte("test")
		r, err := c.Send(ctx, &pb.Transaction{Data: payload})
		if err != nil {
			log.Fatalf("could not send: %v", err)
		}
		log.Printf("Reply: %s", string(r.GetMessage()))
	*/

	contractID := createContract(c, ctx)
	createTransaction(c, ctx, contractID[:])
}

func createContract(c pb.SporeClient, ctx context.Context) [32]byte {
	wasm, err := ioutil.ReadFile("./increment.wasm")
	if err != nil {
		panic(err)
	}
	// create contract transaction
	r, err := c.CreateContract(ctx, &pb.Transaction{Data: wasm, Contract: true})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("Reply: %s", string(hex.EncodeToString(r.GetTransactionId())))

	sum := sha256.Sum256(wasm)
	fmt.Printf("sha256: %x\n", sum)
	return sum
}

func createTransaction(c pb.SporeClient, ctx context.Context, contractID []byte) {
	payload := []byte("increment")
	r, err := c.Send(ctx, &pb.Transaction{Data: payload, To: contractID})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("Reply: %s", string(hex.EncodeToString(r.GetTransactionId())))
}
