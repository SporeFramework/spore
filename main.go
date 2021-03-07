package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	cr "github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"github.com/sporeframework/spore/dag"
	"github.com/sporeframework/spore/protocol"
)

func main2() {
	WasmTime()
	// Gasm()
}

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "sporep2p"

// bootstrappers
type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var bootstrappers arrayFlags

var log = logrus.New()

func main() {
	// parse some flags to set our nickname and the room to join
	flag.Var(&bootstrappers, "connect", "Connect to target bootstrap node. This can be any chat node on the network.")
	listenHost := flag.String("host", "0.0.0.0", "The bootstrap node host listen address")
	port := flag.Int("port", 0, "The node's listening port. This is useful if using this node as a bootstrapper.")
	rpcPort := flag.Int("rpc", 9000, "The node's rpc port.")
	useKey := flag.Bool("use-key", false, "Use an ECSDS keypair as this node's identifier. The keypair is generated if it does not exist in the app's local config directory.")
	info := flag.Bool("info", false, "Display node endpoint information before logging into the main chat room")
	daemon := flag.Bool("daemon", false, "Run as a bootstrap daemon only")
	flag.Parse()

	conf := ConfigSetup()

	ctx := context.Background()

	// Intialize the chain
	dag.InitializeChain()

	var err error
	// DHT Peer routing
	//var idht *dht.IpfsDHT
	routing := libp2p.Routing(func(h host.Host) (cr.PeerRouting, error) {
		return makeDht(ctx, h)
	})

	cm := connmgr.NewConnManager(
		100,         // Lowwater
		400,         // HighWater,
		time.Minute, // GracePeriod
	)

	psk, _ := ClusterSecret()

	var h host.Host
	if *useKey {
		pk := GetKey()
		h, err = libp2p.New(ctx,
			// use a private network
			libp2p.PrivateNetwork(psk),
			// listen addresses
			libp2p.ListenAddrStrings(
				fmt.Sprintf("/ip4/%s/tcp/%d", *listenHost, *port),
			),
			// support TLS connections
			libp2p.Security(libp2ptls.ID, libp2ptls.New),
			// support secio connections
			libp2p.Security(secio.ID, secio.New),
			// support any other default transports (TCP)
			libp2p.DefaultTransports,
			// Let this host use the DHT to find other hosts
			routing,
			// Connection Manager
			libp2p.ConnectionManager(cm),
			// Attempt to open ports using uPNP for NATed hosts.
			libp2p.NATPortMap(),
			// Let this host use relays and advertise itself on relays if
			// it finds it is behind NAT. Use libp2p.Relay(options...) to
			// enable active relays and more.
			libp2p.EnableAutoRelay(),
			// Use the defined identity
			libp2p.Identity(pk),
		)
		LogInfo("🔐 Using identity from key:", h.ID().Pretty())
	} else {
		h, err = libp2p.New(ctx,
			// use a private network
			libp2p.PrivateNetwork(psk),
			// listen addressesß
			libp2p.ListenAddrStrings(
				fmt.Sprintf("/ip4/%s/tcp/%d", *listenHost, *port),
			),
			// support TLS connections
			libp2p.Security(libp2ptls.ID, libp2ptls.New),
			// support secio connections
			libp2p.Security(secio.ID, secio.New),
			// support any other default transports (TCP)
			libp2p.DefaultTransports,
			// Let this host use the DHT to find other hosts
			routing,
			// Connection Manager
			libp2p.ConnectionManager(cm),
			// Attempt to open ports using uPNP for NATed hosts.
			libp2p.NATPortMap(),
			// Let this host use relays and advertise itself on relays if
			// it finds it is behind NAT. Use libp2p.Relay(options...) to
			// enable active relays and more.
			libp2p.EnableAutoRelay(),
		)
	}
	if err != nil {
		log.Error(err)
		panic(err)
	}

	_, err = CollectBootstrapAddrInfos(ctx)

	h.Network().Notify(&network.NotifyBundle{
		ConnectedF: func(n network.Network, c network.Conn) {
			s := fmt.Sprintf("%s/p2p/%s", c.RemoteMultiaddr(), c.RemotePeer())
			if Find(GetConfig().Bootstrappers, s) {
				fmt.Println("🌟 Connected to Bootstrap Node:", s)
			}
		},

		DisconnectedF: func(n network.Network, c network.Conn) {
			s := fmt.Sprintf("%s/p2p/%s", c.RemoteMultiaddr(), c.RemotePeer())
			if Find(GetConfig().Bootstrappers, s) {
				fmt.Println("🛑 Disconnected from Bootstrap Node:", s)

				// thread
				go func(s string, peerId peer.ID) {
					for i := 0; i < 100; i++ {
						fmt.Printf("Loop %s\n", i)
						time.Sleep(2 * time.Second)
						targetAddr, _ := multiaddr.NewMultiaddr(s)
						targetInfo, _ := peer.AddrInfoFromP2pAddr(targetAddr)
						c := h.Network().Connectedness(peerId)
						fmt.Println("Connectedness:", c)

						err = h.Connect(ctx, *targetInfo)
						if err != nil {
							log.Warn("Trying to connect to bootstrap Peer", s, err)
						}
						if h.Network().Connectedness(peerId) == network.Connected {
							return
						}
					}
				}(s, c.RemotePeer())

			}
		},
	})

	fmt.Println("🌟 Id:", h.ID().Pretty())
	// print the node's listening addresses
	fmt.Println("🔖 Listen addresses:", h.Addrs())

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	sub, err := ps.Subscribe(PubsubTopic)
	if err != nil {
		panic(err)
	}
	go pubsubHandler(ctx, sub)

	// setup local mDNS discovery
	err = setupMdnsDiscovery(ctx, h)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	donec := make(chan struct{}, 1)
	//go chatInputLoop(ctx, h, ps, donec)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)

	if *info {
		fmt.Println("🔖  Network id:", conf.ClusterKey)
		fmt.Print("👢 Available endpoints: \n")
		for _, addr := range h.Addrs() {
			fmt.Printf("	%s/p2p/%s\n", addr, h.ID().Pretty())
			log.Info("	%s/p2p/%s\n", addr, h.ID().Pretty())
		}
		fmt.Println("Press any key to continue...")
		fmt.Scanln() // wait for Enter Key
	}

	go protocol.StartRPCServer(PubsubTopic, ps, rpcPort)

	if *daemon {
		// select {}
		// TODO remove this
		select {
		case <-stop:
			h.Close()
			os.Exit(0)
		case <-donec:
			h.Close()
		}
	} else {

		select {
		case <-stop:
			h.Close()
			os.Exit(0)
		case <-donec:
			h.Close()
		}
		// draw the UI
		/*
			ui := NewChatUI(cr)
			if err = ui.Run(); err != nil {
				printErr("error running text UI: %s", err)
				log.Error("error running text UI: %s", err)
			}
		*/
	}
}

func makeDht(ctx context.Context, h host.Host) (*dht.IpfsDHT, error) {
	dht.DefaultBootstrapPeers = nil
	bootstrapPeers, err := CollectBootstrapAddrInfos(ctx)
	idht, _ := dht.New(ctx, h,
		dht.Mode(dht.ModeServer),
		dht.ProtocolPrefix("/sporep2p/kad/1.0.0"),
		dht.BootstrapPeers(bootstrapPeers...),
	)

	fmt.Println("Bootstrapping the DHT")
	if err = idht.Bootstrap(ctx); err != nil {
		panic(err)
	}
	return idht, err
}

// printErr is like fmt.Printf, but writes to stderr.
func printErr(m string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, m, args...)
}

// defaultNick generates a nickname based on the $USER environment variable and
// the last 8 chars of a peer ID.
func defaultNick(p peer.ID) string {
	return fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(p))
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	log.Info("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		log.Error("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupMdnsDiscovery(ctx context.Context, h host.Host) error {
	// setup mDNS discovery to find local peers
	disc, err := discovery.NewMdnsService(ctx, h, DiscoveryInterval, DiscoveryServiceTag)
	if err != nil {
		return err
	}

	n := discoveryNotifee{h: h}
	disc.RegisterNotifee(&n)
	return nil
}
