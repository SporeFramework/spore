package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	pb "github.com/sporeframework/spore/protocol"
)

var handles = map[string]string{}

const PubsubTopic = "/libp2p/spore/chat/1.0.0"

func pubsubCreateContractHandler(id peer.ID, txn *pb.Transaction) {
	contractID, gas, err := CreateWasmContract(txn.Data)
	if err != nil {
		fmt.Printf("An error has occured: %s\n", err.Error())
	}
	fmt.Printf("Contract created, ID: %s, Gas: %s\n", hex.EncodeToString(contractID[:]), gas)
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
	/*
		handle, ok := handles[id.String()]
		if !ok {
			handle = id.ShortString()
		}
		fmt.Printf("%s: %s\n", handle, msg.Data)
	*/
}

func pubsubUpdateHandler(id peer.ID, msg *pb.UpdatePeer) {
	oldHandle, ok := handles[id.String()]
	if !ok {
		oldHandle = id.ShortString()
	}
	handles[id.String()] = string(msg.UserHandle)
	fmt.Printf("%s -> %s\n", oldHandle, msg.UserHandle)
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
		case pb.Request_UPDATE_PEER:
			pubsubUpdateHandler(msg.GetFrom(), req.UpdatePeer)
		}
	}
}
