# Spore 🍄
Spore is a layer 2 framework for building independently decentralized apps

# Test Network

A simple test network can be constructed with docker/docker-compose.

1. Build the base docker image: 
    `docker build -t spore-node .`
2. Bring up the network, scaling up to however many nodes you wish: 
    `docker-compose up --build --scale spore-node=6`


# Wasm Contracts

Smart contracts on Spore are deployed as WebAssembly binaries, opening up Dapp development to several languages including, Go, Rust, Solidity, AssemblyScript, c, and a plethora of others.

## Wasm precompiled functions

### Cryptography
* sha256 - a sha3 hash

### Bridging
* sendBitcoinTxn - Send a Bitcoin signed transaction to the Bitcoin blockchain. Returns the transaction. Msg.sender must send a signed raw transaction, along with metadata including to, from, amount (in Satoshis)
* sendEthereumTxn - Send a Ethereum signed transaction to the Ethereum blockchain. This can include a simple value transfer, ERC20 transfer, or any interaction with 

### Oracles
* get - HTTP get content
