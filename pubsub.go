package main

import (
	"context"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	pb "github.com/sporeframework/spore/protocol"
)

var handles = map[string]string{}

const PubsubTopic = "/libp2p/spore/chat/1.0.0"

func pubsubTransactionHandler(id peer.ID, msg *pb.SendTransaction) {
	handle, ok := handles[id.String()]
	if !ok {
		handle = id.ShortString()
	}
	fmt.Printf("%s: %s\n", handle, msg.Data)
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
			pubsubTransactionHandler(msg.GetFrom(), req.SendTransaction)
		case pb.Request_UPDATE_PEER:
			pubsubUpdateHandler(msg.GetFrom(), req.UpdatePeer)
		}
	}
}
