package decoder

import (
	"time"

	"github.com/bruxaodev/sol-listner/pkg/config"
	"github.com/bruxaodev/sol-listner/pkg/decoder/pumpfun"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

type Decoder int

const (
	Pumpfun Decoder = iota
	Block
)

func BlockDecoder(data []byte, decoder Decoder) {

	switch decoder {
	case Pumpfun:
		now := time.Now()
		var subscribeData pumpfun.SubscribeData
		err := json.Unmarshal(data, &subscribeData)
		if err != nil {
			config.Logger.Error("decode block", zap.Error(err))
		}

		config.Logger.Info("decode block", zap.Int("slot", subscribeData.Params.Result.Context.Slot), zap.Duration("time", time.Since(now)))
		go NewBlock(subscribeData)
		go NewToken(subscribeData)
	case Block:
		config.Logger.Info("decode block", zap.String("data", string(data)))
	}
}
