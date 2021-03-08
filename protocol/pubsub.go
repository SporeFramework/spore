package protocol

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"sync"

	"github.com/kirsle/configdir"
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	log "github.com/sirupsen/logrus"
	"github.com/sporeframework/spore/contract"
	dag "github.com/sporeframework/spore/dag"
	"github.com/sporeframework/spore/db"
	"google.golang.org/protobuf/proto"
)

const (
	PubsubTopic       = "/spore/1.0.0"
	DatabaseNamespace = "sporedb"
)

var (
	g        *dag.GreedyGraphMem
	Database db.DB
)

func InitializeChain() {
	var k = 1621
	graph, err := dag.NewGreedyGraphMem(k)
	if err != nil {
		log.Error("failed to create new GreedyGraphMem: %s", err)
	}
	g = graph

	// startup the db
	// BadgerDB database
	dbPath := configdir.LocalConfig("spore", "db")
	err = configdir.MakePath(dbPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}
	Database, err = db.NewBadgerDB(dbPath)
	if err != nil {
		panic(err)
	}
}

var mu sync.Mutex

func AddBlock(txn *Transaction) {

	Name := txn.Id
	// this is only really required when using ordering
	mu.Lock()
	defer mu.Unlock()

	tips, err := g.Tips()
	if err != nil {
		fmt.Errorf("❌ failed to get tips", Name, err)
	}

	ok, err := g.Add(string(Name), tips)
	if err != nil {
		fmt.Errorf("❌ failed to add node %s: %s", Name, err)
	}

	if !ok {
		fmt.Errorf("❌ node %s not added to graph", Name)
	}

	// write transaction to the database
	go set(txn)

	// debug
	/*
		getTxnBytes, err := database.Get([]byte(DatabaseNamespace), txn.Id)
		if err != nil {
			fmt.Errorf("❌ failed to get node %s from db: %s", Name, err)
		}

		txnUnmarsh := &pb.Transaction{}
		proto.Unmarshal(getTxnBytes, txnUnmarsh)
		fmt.Printf("Got transaction: ", txnUnmarsh)

		nodeSize := len(g.Nodes())
		fmt.Println("Node count: ", nodeSize)
		if nodeSize%100 == 0 {
			fmt.Println("Ordering started...")

			ordered, err := g.Order()
			if err != nil {
				fmt.Errorf("❌ failed to order nodes after adding node %s: %s", Name, err)
			}

			//fmt.Println("Ordering completed: ", ordered)
			fmt.Println("Ordering completed: ", ordered[len(ordered)-5:])
			fmt.Println("Order count: ", len(ordered))
		}
	*/

	fmt.Println("Node count: ", len(g.Nodes()))

}

func set(txn *Transaction) {
	// add to the database
	txnBytes, err := proto.Marshal(txn)
	if err != nil {
		fmt.Errorf("❌ failed to add node %s to db: %s", txn.Id, err)
	}
	go Database.Set([]byte(DatabaseNamespace), txn.Id, txnBytes)
	fmt.Println("Inserted key into db: ", hex.EncodeToString(txn.Id))
}

func createContractHandler(id peer.ID, txn *Transaction) {
	contractID, gas, err := contract.CreateWasmContract(txn.Data)
	if err != nil {
		fmt.Printf("An error has occured: %s\n", err.Error())
	}
	fmt.Printf("Contract created, ID: %s, Gas: %s\n", hex.EncodeToString(contractID[:]), gas)

	AddBlock(txn)
}

func transactionHandler(id peer.ID, txn *Transaction) {
	var contractID [32]byte
	copy(contractID[:], txn.To[:32])

	fmt.Printf("Calling Contract ID: %s\n", hex.EncodeToString(contractID[:]))

	result, gas, err := contract.Call(contractID, string(txn.Data))
	if err != nil {
		fmt.Printf("An error has occured: %s\n", err.Error())
	}
	fmt.Printf("result: %s, gas: %s\n", result, gas)
	AddBlock(txn)
}

func PubsubHandler(ctx context.Context, sub *pubsub.Subscription) {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		req := &Request{}
		err = proto.Unmarshal(msg.Data, req)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		switch req.Type {
		case Request_SEND_TRANSACTION:
			transactionHandler(msg.GetFrom(), req.Transaction)
		case Request_CREATE_CONTRACT:
			createContractHandler(msg.GetFrom(), req.Transaction)
		}
	}
}
