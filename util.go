package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kirsle/configdir"
	"github.com/libp2p/go-libp2p-core/peer"
	crypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/multiformats/go-multiaddr"
)

const defaultConfig = `{
    "ClusterKey": "f73792a8ba5fa5306039ccd82f79887b3319457752ff0b604fc736c72134e336",
    "Bootstrappers": [
		"/ip4/35.224.203.143/tcp/4001/p2p/QmfNdsi6tQfuQ1AbiVbTwxziaRCzuamjP711y42mNW33DS"
	]
}`

// Configuration is the deserialized version of the json configuration
type Configuration struct {
	ClusterKey    string
	Bootstrappers []string
}

// GetConfig loads the configuration json file
func GetConfig() *Configuration {
	configFile := configdir.LocalConfig("spore", "conf.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err := ioutil.WriteFile(configFile, []byte(defaultConfig), 0644)
		if err != nil {
			log.Error(err)
			panic(err)
		}
	}

	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Bootstrappers)
	return &configuration
}

// ConfigSetup sets up the configuration directory
func ConfigSetup() *Configuration {
	// Ensure config directory exists
	configPath := configdir.LocalConfig("spore")
	er := configdir.MakePath(configPath) // Ensure it exists.
	if er != nil {
		panic(er)
	}

	// set up logging
	logfile := configdir.LocalConfig("spore", "spore.log")
	file, erro := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if erro == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	fmt.Println(configPath)
	// keyfile
	keyfile := configdir.LocalConfig("spore", ".key")
	if _, err := os.Stat(keyfile); os.IsNotExist(err) {
		createKey()
	}

	return GetConfig()
}

// GetKey gets the ECDSA private key from the config
// directory
func GetKey() crypto.PrivKey {
	keyfile := configdir.LocalConfig("spore", ".key")

	// open private key file
	content, err := ioutil.ReadFile(keyfile)
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
	return privNew
}

// CreateKey generates a new ECDSA Key Pair and stores it
// in the config directory
func createKey() {
	keyfile := configdir.LocalConfig("spore", ".key")
	// Create a new ECDSA key pair for this host.
	prvKey, _, err := crypto.GenerateECDSAKeyPair(rand.Reader)
	if err != nil {
		panic(err)
	}
	privB, err := prvKey.Bytes()
	if err != nil {
		panic(err)
	}
	pk := fmt.Sprintf("%x", string(privB))
	ioutil.WriteFile(keyfile, []byte(pk), 0644)
	fmt.Println("🔑 ECDSA key generated")
}

// ClusterSecret parses the hex-encoded secret string, checks that it is exactly
// 32 bytes long and returns its value as a byte-slice.x
func ClusterSecret() ([]byte, error) {
	secret, err := hex.DecodeString(GetConfig().ClusterKey)
	if err != nil {
		return nil, err
	}
	switch secretLen := len(secret); secretLen {
	case 0:
		log.Warning("Cluster secret is empty, cluster will start on unprotected network.")
		return nil, nil
	case 32:
		return secret, nil
	default:
		return nil, fmt.Errorf("input secret is %d bytes, cluster secret should be 32", secretLen)
	}
}

// CollectBootstrapAddrInfos converts bootstrap address in config directory
// to a slice of []peer.AddrInfo
func CollectBootstrapAddrInfos(ctx context.Context) ([]peer.AddrInfo, error) {
	bootstrappers := GetConfig().Bootstrappers

	if len(bootstrappers) == 0 {
		LogInfo("🔔 No bootstrappers defined for this node.")
	}

	addrInfoSlice := make([]peer.AddrInfo, len(bootstrappers))
	for i, s := range bootstrappers {
		targetAddr, err := multiaddr.NewMultiaddr(s)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		targetInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		addrInfoSlice[i] = *targetInfo
		LogInfo("🔔 Calling bootstrap node:", s)
	}

	return addrInfoSlice, nil
}

// LogInfo logs to console and logger
func LogInfo(m string, args ...interface{}) {
	fmt.Println(m, args)
	log.Info(m, args)
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
