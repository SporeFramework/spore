package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	dag "github.com/sporeframework/spore/dag"
	pb "github.com/sporeframework/spore/protocol"
)

const PubsubTopic = "/spore/1.0.0"

var mu sync.Mutex

func AddBlock(Name string) *dag.Block {

	mu.Lock()
	defer mu.Unlock()
	// block := dag.InsertBlock(strconv.Itoa(rand.Int()))
	block := dag.InsertBlock(Name)

	/*
		block := dag.Block{Name, -1, -1, make(map[string]*dag.Block), make(map[string]*dag.Block), make(map[string]bool)}
		for _, ref := range keys {
			block.Prev[ref] = chain[ref]
			chain[ref].Next[Name] = &block
		}
		mu.Lock()
		block.SizeOfPastSet = dag.SizeOfPastSet(&block)
		mu.Unlock()
		chain[Name] = &block
	*/

	fmt.Println("Added Block: ", block)

	// print out the entire chain for debugging purposes
	//debugChain()

	//block := dag.Block{Name, -1, -1, make(map[string]*dag.Block), make(map[string]*dag.Block), make(map[string]bool)}
	//chain[Name] = &block

	//tips = dag.FindTips(chain)
	//fmt.Println("tips: ", tips)
	//tipsName := dag.LTPQ(chain, true) // LTPQ is not relevant here, I just use it to get Tips name.
	//ChainAddBlock("Virtual", []string{Name})

	return block
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
