package main

import (
	"solana-listner/pkg/connection"
	"solana-listner/pkg/decoder"
	"time"
)

func main() {
	pumpChannel := make(chan []byte)
	rpc := connection.NewRpc()

	go rpc.Connection(
		"ws://172.82.90.165:8900",
		connection.SUBSCRIBE_PUMPFUN,
		pumpChannel,
	)

	go func() {
		for {
			message := <-pumpChannel
			decoder.BlockDecoder(message, decoder.Pumpfun)
		}
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}