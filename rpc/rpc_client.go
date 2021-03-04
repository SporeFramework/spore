package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/valyala/gorpc"
)

func main() {
	c := &gorpc.Client{
		// TCP address of the server.
		Addr: "127.0.0.1:9000",
	}
	c.Start()

	payload := "abc123"
	// All client methods issuing RPCs are thread-safe and goroutine-safe,
	// i.e. it is safe to call them from multiple concurrently running goroutines.
	resp, err := c.Call(payload)
	if err != nil {
		log.Fatalf("Error when sending request to server: %s", err)
	}
	if resp.(string) != payload {
		log.Fatalf("Unexpected response from the server: %+v", resp)
	}
}
