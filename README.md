# Spore
Spore is a layer 2 framework for building independently decentralized apps

# Test Network

A simple test network can be constructed with docker/docker-compose.

1. Build the base docker image: 
    `docker build -t spore-node .`
2. Bring up the network, scaling up to however many nodes you wish: 
    `docker-compose up --build --scale spore-node=6`
