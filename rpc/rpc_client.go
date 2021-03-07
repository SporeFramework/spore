package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	pb "github.com/sporeframework/spore/protocol"
	"golang.org/x/crypto/sha3"
	"google.golang.org/grpc"

	"math/rand"
)

func generateRandomKey() (address []byte, privateKey *ecdsa.PrivateKey) {
	privateKey, _ = crypto.GenerateKey()

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("Private key: %s\n", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("Public key:\t %s\n", hexutil.Encode(publicKeyBytes)[4:])

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Bytes()
	fmt.Printf("Public address (from ECDSA): \t%s\n", crypto.PubkeyToAddress(*publicKeyECDSA).Hex())

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Printf("Public address (Hash of public key):\t%s\n", hexutil.Encode(hash.Sum(nil)[12:]))

	return address, privateKey
}

func main() {

	rpcPort := flag.Int("rpc", 9000, "The node's rpc port.")

	addr := "localhost:" + strconv.Itoa(*rpcPort)
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSporeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	address, privateKey := generateRandomKey()

	contractID := createContract(c, ctx, address, privateKey)

	// contractID, _ := hex.DecodeString("f5b012bbab7f165bc5eec4302f0952c1f3f8b601d4907cb8c5ae781a71821abb")

	for i := 0; i < 300; i++ {
		rand.Int31()
		createTransaction(c, ctx, contractID[:], address, privateKey)
		/*
			if i%1000 == 0 {
				time.Sleep(2 * time.Second)
			}
		*/
		fmt.Println(i)
	}

}

func createContract(c pb.SporeClient, ctx context.Context, address []byte, prv *ecdsa.PrivateKey) [32]byte {
	wasm, err := ioutil.ReadFile("./increment.wasm")
	if err != nil {
		panic(err)
	}

	// create the transaction
	txn := &pb.Transaction{
		Data:     wasm,
		From:     address,
		Contract: true,
		Nonce:    rand.Int31(),
	}

	// sign the transaction
	txnBytes, _ := proto.Marshal(txn)
	pSum := sha256.Sum256(txnBytes)
	fmt.Printf("pSum: %s\n", hex.EncodeToString(pSum[:]))
	sig, err := crypto.Sign(pSum[:], prv)
	txn.Signature = sig
	fmt.Printf("sig: %s\n", hex.EncodeToString(sig))

	// create contract transaction
	r, err := c.CreateContract(ctx, txn)
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("Reply: %s", string(hex.EncodeToString(r.GetTransactionId())))

	sum := sha256.Sum256(wasm)
	fmt.Printf("sha256: %x\n", sum)
	return sum
}

func createTransaction(c pb.SporeClient, ctx context.Context, contractID []byte, address []byte, prv *ecdsa.PrivateKey) {
	payload := []byte("increment")

	// create the transaction
	txn := &pb.Transaction{
		Data:     payload,
		To:       contractID,
		From:     address,
		Contract: true,
		Nonce:    rand.Int31(),
	}
	// sign the transaction
	txnBytes, _ := proto.Marshal(txn)
	pSum := sha256.Sum256(txnBytes)
	fmt.Printf("pSum: %s\n", hex.EncodeToString(pSum[:]))
	sig, err := crypto.Sign(pSum[:], prv)
	txn.Signature = sig
	//fmt.Printf("sig: %s\n", hex.EncodeToString(sig))

	txnBytes, _ = proto.Marshal(txn)
	txnHash := sha256.Sum256(txnBytes)
	fmt.Printf("txnId: %s\n", hex.EncodeToString(txnHash[:]))

	r, err := c.Send(ctx, txn)
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("Reply: %s", string(hex.EncodeToString(r.GetTransactionId())))
}
