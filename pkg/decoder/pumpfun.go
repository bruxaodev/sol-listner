package decoder

import (
	"log"
	"solana-listner/pkg/decoder/pumpfun"
)

func NewToken(data pumpfun.SubscribeData) {
	for _, transaction := range data.Params.Result.Value.Block.Transactions {
		go func() {
			instructions := transaction.Transaction.Message.Instructions
			for _, instruction := range instructions {
				if len(instruction.Accounts) >= 2 && instruction.Accounts[1] == "TSLvdd1pWpHVjahSpsvCXUbgwsL3JAcvokwaKt1eokM" {
					log.Println("New Token: ", instruction.Accounts[0])
				}
			}
		}()
	}
}
