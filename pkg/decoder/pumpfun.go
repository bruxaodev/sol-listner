package decoder

import (
	"github.com/bruxaodev/sol-listner/pkg/config"
	"github.com/bruxaodev/sol-listner/pkg/decoder/pumpfun"
	"go.uber.org/zap"
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
							config.Logger.Info("new token", zap.Int("block", data.Params.Result.Context.Slot), zap.String("tx", transaction.Transaction.Signatures[0]))
							continue
						}
					}
				}
			}
		}()
	}
}

func NewBlock(data pumpfun.SubscribeData) {
	config.Logger.Info("block", zap.Int("slot", data.Params.Result.Context.Slot), zap.Int("totalTransactions", len(data.Params.Result.Value.Block.Transactions)))
}
