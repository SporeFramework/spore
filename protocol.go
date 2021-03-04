package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"time"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/valyala/gorpc"
)

func sendTransaction(ps *pubsub.PubSub, msg string) {
	msgID := make([]byte, 10)
	_, err := rand.Read(msgID)
	defer func() {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()
	if err != nil {
		return
	}
	now := time.Now().Unix()
	req := &Request{
		Type: Request_SEND_TRANSACTION.Enum(),
		SendTransaction: &SendTransaction{
			Id:      msgID,
			Data:    []byte(msg),
			Created: &now,
		},
	}
	msgBytes, err := req.Marshal()
	if err != nil {
		return
	}
	err = ps.Publish(pubsubTopic, msgBytes)
}

func updatePeer(ps *pubsub.PubSub, id peer.ID, handle string) {
	oldHandle, ok := handles[id.String()]
	if !ok {
		oldHandle = id.ShortString()
	}
	handles[id.String()] = handle

	req := &Request{
		Type: Request_UPDATE_PEER.Enum(),
		UpdatePeer: &UpdatePeer{
			UserHandle: []byte(handle),
		},
	}
	reqBytes, err := req.Marshal()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = ps.Publish(pubsubTopic, reqBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("%s -> %s\n", oldHandle, handle)
}

func startRPCServer(ps *pubsub.PubSub, port *int) {

	s := &gorpc.Server{
		// Accept clients on this TCP address.
		Addr: ":" + strconv.Itoa(*port),

		// Echo handler - just return back the message we received from the client
		Handler: func(clientAddr string, request interface{}) interface{} {
			fmt.Printf("Obtained request %+v from the client %s\n", request, clientAddr)
			sendTransaction(ps, fmt.Sprintf("%v", request))
			return request
		},
	}
	if err := s.Serve(); err != nil {
		fmt.Printf("Cannot start rpc server: %s", err)
	}

}

/*
func chatInputLoop(ctx context.Context, h host.Host, ps *pubsub.PubSub, donec chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if strings.HasPrefix(msg, "/name ") {
			newHandle := strings.TrimPrefix(msg, "/name ")
			newHandle = strings.TrimSpace(newHandle)
			updatePeer(ps, h.ID(), newHandle)
		} else {
			sendMessage(ps, msg)
		}
	}
	donec <- struct{}{}
}
*/
