package decoder

import (
	"log"

	"github.com/bruxaodev/sol-listner/pkg/decoder/pumpfun"
)

var (
	PUMP_PROGRAM_ID = "6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P"
	PUMP_MINT       = "TSLvdd1pWpHVjahSpsvCXUbgwsL3JAcvokwaKt1eokM"
)

func NewToken(data pumpfun.SubscribeData) {
	for _, transaction := range data.Params.Result.Value.Block.Transactions {
		go func() {
			instructions := transaction.Transaction.Message.Instructions
			for _, instruction := range instructions {
				if instruction.ProgramId == PUMP_PROGRAM_ID {
					for _, account := range instruction.Accounts {
						if account == PUMP_MINT {
							log.Printf("new token - block: %v  tx: https://solscan.io/tx/%s \n", data.Params.Result.Context.Slot, transaction.Transaction.Signatures[0])
							continue
						}
					}
				}
			}
		}()
	}
}

func NewBlock(data pumpfun.SubscribeData) {
	log.Printf("block: %d totalTransactions: %d\n", data.Params.Result.Context.Slot, len(data.Params.Result.Value.Block.Transactions))
}
