package main

import (
	"os"
	"time"

	"github.com/bruxaodev/sol-listner/pkg/connection"
	"github.com/bruxaodev/sol-listner/pkg/decoder"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	blockDataChannel := make(chan []byte)
	defer close(blockDataChannel)
	rpc := connection.NewRpc()

	go rpc.Connection(
		os.Getenv("WS_RPC"),
		connection.SUBSCRIBE_PUMPFUN,
		blockDataChannel,
	)

	go func() {
		for {
			blockData := <-blockDataChannel
			go func(message []byte) {
				decoder.BlockDecoder(message, decoder.Pumpfun)
			}(blockData)
		}
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
