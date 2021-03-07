package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	log "github.com/sirupsen/logrus"
	dag "github.com/sporeframework/spore/dag"
	pb "github.com/sporeframework/spore/protocol"
)

const PubsubTopic = "/spore/1.0.0"

var g *dag.GreedyGraphMem

func InitializeChain() {
	var k = 1000
	graph, err := dag.NewGreedyGraphMem(k)
	if err != nil {
		log.Error("failed to create new GreedyGraphMem: %s", err)
	}
	g = graph
}

func AddBlock(Name string) {

	ok, err := g.AddNode(Name)
	if err != nil {
		log.Error("failed to add node %s: %s", Name, err)
	}

	if !ok {
		log.Error("node %s not added to graph", Name)
	}

	// debug
	/*
		nodeSize := len(g.Nodes())
		fmt.Println("Node count: ", nodeSize)

		if nodeSize%100 == 0 {
			fmt.Println("Ordering started...")

			ordered, err := g.Order()
			if err != nil {
				fmt.Errorf("failed to order nodes after adding node %s: %s", Name, err)
			}

			fmt.Println("Ordering completed: ", ordered[len(ordered)-5:])
			fmt.Println("Ordering completed.")
		}
	*/
}

func pubsubCreateContractHandler(id peer.ID, txn *pb.Transaction) {
	contractID, gas, err := CreateWasmContract(txn.Data)
	if err != nil {
		fmt.Printf("An error has occured: %s\n", err.Error())
	}
	fmt.Printf("Contract created, ID: %s, Gas: %s\n", hex.EncodeToString(contractID[:]), gas)

	AddBlock(hex.EncodeToString(txn.Id))
}

func pubsubTransactionHandler(id peer.ID, txn *pb.Transaction) {

	var contractID [32]byte
	copy(contractID[:], txn.To[:32])

	fmt.Printf("Calling Contract ID: %s\n", hex.EncodeToString(contractID[:]))

	result, gas, err := Call(contractID, string(txn.Data))
	if err != nil {
		fmt.Printf("An error has occured: %s\n", err.Error())
	}
	fmt.Printf("result: %s, gas: %s\n", result, gas)
	AddBlock(hex.EncodeToString(txn.Id))
}

func pubsubHandler(ctx context.Context, sub *pubsub.Subscription) {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		req := &pb.Request{}
		err = proto.Unmarshal(msg.Data, req)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		switch req.Type {
		case pb.Request_SEND_TRANSACTION:
			pubsubTransactionHandler(msg.GetFrom(), req.Transaction)
		case pb.Request_CREATE_CONTRACT:
			pubsubCreateContractHandler(msg.GetFrom(), req.Transaction)
		}
	}
}
