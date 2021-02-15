package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	// "crypto/rand"
	"os"
	"io/ioutil"
	"encoding/hex"
	"strings"
)

func main() {
	help := flag.Bool("help", false, "Display Help")
	listenHost := flag.String("host", "0.0.0.0", "The bootstrap node host listen address\n")
	port := flag.Int("port", 4001, "The bootstrap node listen port")
	flag.Parse()

	if *help {
		fmt.Printf("This is a simple bootstrap node for kad-dht application using libp2p\n\n")
		fmt.Printf("Usage: \n   Run './bootnode'\nor Run './bootnode -host [host] -port [port]'\n")

		os.Exit(0)
	}

	fmt.Printf("[*] Listening on: %s with port: %d\n", *listenHost, *port)

	ctx := context.Background()
	//r := mrand.New(mrand.NewSource(int64(*port)))

	// Creates a new ECDSA key pair for this host.
	/*
	prvKey, _, err := crypto.GenerateECDSAKeyPair(rand.Reader)
	if err != nil {
		panic(err)
	}

	privB, err := prvKey.Bytes()
	if err != nil {
		panic(err)
	}

	fmt.Printf("private key: %x \n", string(privB))
	*/

    content, err := ioutil.ReadFile("key.txt")
    if err != nil {
        panic(err)
    }

    hexString := strings.TrimSuffix(string(content), "\n")
    decoded, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	privNew, err := crypto.UnmarshalPrivateKey(decoded)
	if err != nil {
		panic(err)
	}

	// 0.0.0.0 will listen on any interface device.
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", *listenHost, *port))

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	host, err := libp2p.New(
		ctx,
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(privNew),
	)

	if err != nil {
		panic(err)
	}

	_, err = dht.New(ctx, host)
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	fmt.Printf("[*] Your Bootstrap ID Is: /ip4/%s/tcp/%v/p2p/%s\n", *listenHost, *port, host.ID().Pretty())
	fmt.Println("")
	select {}
}
